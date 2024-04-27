package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/happyin3/buy/product/dao"
	"github.com/happyin3/buy/product/service"
)

func main() {
	dao.ConnectDatabase()

	product := service.NewProductService()

	app := fiber.New()

	app.Get("/products", product.Products)
	app.Get("/products/:id", product.Get)
	app.Post("/products", product.Create)
	app.Put("/products/:id", product.Update)
	app.Delete("/products/:id", product.Delete)

	app.Listen(":3000")
}
