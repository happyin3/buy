package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/happyin3/buy/user/dao"
	"github.com/happyin3/buy/user/service"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	dao.ConnectDatabase()

	user := service.NewUserService()

	app := fiber.New()

	// 定义一个中间件来处理响应格式
	//app.Use(func(c fiber.Ctx) error {
	//	// 处理请求
	//	err := c.Next()

	//	// 处理响应
	//	if err != nil {
	//		return c.JSON(Response{
	//			Success: false,
	//			Message: err.Error(),
	//			Data:    nil,
	//		})
	//	}

	//	return c.JSON(Response{
	//		Success: true,
	//		Message: "Success",
	//		Data:    c.Response().Body(),
	//	})
	//})

	app.Get("/users", user.List)
	app.Get("/users/:id", user.Get)
	app.Post("/users", user.Create)
	app.Put("/users/:id", user.Update)
	app.Delete("/users/:id", user.Delete)

	app.Listen(":3001")
}
