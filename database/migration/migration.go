package migration

import (
	"restapi-gorm-gofiber/database"
	"restapi-gorm-gofiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Users{}, &entity.Books{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Migration successful")
	
}
