package main

import (
	"restapi-gorm-gofiber/database"
	"restapi-gorm-gofiber/database/migration"
	"restapi-gorm-gofiber/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}
