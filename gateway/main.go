package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Post("/login", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "success",
			"data":    nil,
		})
	})

	app.Listen(":3000")
}
