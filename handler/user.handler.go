package handler

import (
	"restapi-gorm-gofiber/database"
	"restapi-gorm-gofiber/model/entity"
	"restapi-gorm-gofiber/model/request"
	// "restapi-gorm-gofiber/model/response"
	"restapi-gorm-gofiber/utils"


	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"


	"log"
)


func UserHandlerGetAll(ctx *fiber.Ctx) error {
	userInfo := ctx.Locals("userInfo")
	log.Println(userInfo)
	

	var users []entity.Users
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)


}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// Request Validation
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Validation error",
			"error": errValidate.Error(),
		})
	}

	newUser := entity.Users{
		Name: user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}	

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser!= nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Could not create user", 
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully created",
		"data": newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.Users
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//userResponse := response.UserResponse{
	//	ID:        user.ID,
	//	Name:      user.Name,
	//	Email:     user.Email,
	//	Address:   user.Address,
	//	Phone:     user.Phone,
	//	CreatedAt: user.CreatedAt,
	//	UpdatedAt: user.UpdatedAt,
	//}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	userId := ctx.Params("id")

	var user entity.Users

	err := database.DB.First(&user, "10", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"data": nil,
		})
	}




	// Update user
	user.Name = userRequest.Name
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Phone = userRequest.Phone

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not update user",
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully created",
		"data": user,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateEmailRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	userId := ctx.Params("id")

	var user entity.Users

	errCheckEmail := database.DB.First(&user, "10", userId).Error
	if errCheckEmail != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Email already use",
			"data": nil,
		})
	}




	// Update email
	user.Email = userRequest.Email

	

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not update user",
		})
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "Email successfully updated",
		"data": user,
	})
}


func UserHandlerDelete(ctx *fiber.Ctx) error {
	// get all users by id
	userId := ctx.Params("id")

	var user entity.Users
	err := database.DB.First(&user, "10", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"data": nil,
		})
	}

	// Delete user, 
	errDelete := database.DB.Delete(&user).Error
	// if failed
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not delete user",
		})
	}
	// if success
	return ctx.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully deleted",
		"data": nil,
	})

	
}
