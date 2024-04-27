package service

import (
	"github.com/gofiber/fiber/v3"
	"github.com/happyin3/buy/product/dao"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) Products(c fiber.Ctx) error {
	var datas []dao.Product
	dao.Database.Find(&datas)
	return c.JSON(datas)
}

func (p *ProductService) Get(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")

	var data dao.Product
	dao.Database.First(&data, id)
	return c.JSON(data)
}

func (p *ProductService) Create(c fiber.Ctx) error {
	data := new(dao.Product)

	if err := c.Bind().JSON(data); err != nil {
		return err
	}

	dao.Database.Create(&data)

	return c.JSON(data)
}

func (p *ProductService) Update(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")

	data := new(dao.Product)

	if err := c.Bind().JSON(data); err != nil {
		return err
	}

	m := new(dao.Product)
	dao.Database.First(&m, id)
	dao.Database.Model(&m).Updates(data)

	return c.JSON(m)
}

func (p *ProductService) Delete(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")

	dao.Database.Delete(&dao.Product{}, id)

	return c.JSON(map[string]int{"code": 200})
}
