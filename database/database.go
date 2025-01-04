package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	const MYSQL = "root:BackSpace@tcp(127.0.0.1:3306)/restapi_gorm_gofiber?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connected!")
	}
}
