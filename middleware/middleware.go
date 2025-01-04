package middleware

import (
    "github.com/gofiber/fiber/v2"
	"restapi-gorm-gofiber/utils"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthoticated",
		})
	}


	// _, err := utils.VerifyToken(token) 
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthoticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "forbidden access",
			
		})
	} 

	ctx.Locals("userInfo", claims)

	return ctx.Next()
}