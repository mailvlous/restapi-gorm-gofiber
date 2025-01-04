package handler

import (
	"fmt"
	"log"
	"restapi-gorm-gofiber/database"
	"restapi-gorm-gofiber/model/entity"
	"restapi-gorm-gofiber/model/request"

	// "restapi-gorm-gofiber/model/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)

	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	// Request Validation
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Validation error",
			"error": errValidate.Error(),
		})
	}

	// HANDLE FILE
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to get file",
			"data": nil,
		})
	}

	var filename string

	if file != nil {

		filename := file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
		if errSaveFile!= nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "Failed to save file",
				"data": nil,
			})
		}
	} else {
		log.Println("Nothing to be saved")
	}



	newBook := entity.Books{
		Title: book.Title,
		Author: book.Author,
		Cover: filename,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook!= nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed to store data", 
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data": newBook,
	})


}