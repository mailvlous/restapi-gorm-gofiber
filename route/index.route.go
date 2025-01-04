package route

import (
	"restapi-gorm-gofiber/config"
	"restapi-gorm-gofiber/handler"
	"restapi-gorm-gofiber/middleware"
	"github.com/gofiber/fiber/v2"
)


func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/assets")

	r.Post ("/login", handler.LoginHandler)

	r.Get ("/user", middleware.Auth, handler.UserHandlerGetAll)
	// r.Get ("/user", handler.UserHandlerGetAll)
	r.Get ("/user/:id", handler.UserHandlerGetById)
	r.Post ("/user", handler.UserHandlerCreate)
	r.Put ("/user/:id", handler.UserHandlerUpdate)
	r.Put ("/user/:id/updateEmail", handler.UserHandlerUpdateEmail)
	r.Delete ("/user/:id", handler.UserHandlerDelete)

	//BOOK

	r.Post("/book", handler.BookHandlerCreate)
	




	
}
