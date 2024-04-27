package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/gofiber/fiber/v3"
	util "github.com/happyin3/buy/uilts/utils"
	"github.com/happyin3/buy/user/dao"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetByFilter(c fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.Bind().JSON(data); err != nil {
		return err
	}

	var user dao.User
	dao.Database.Where(data).First(&user)

	res := util.FmtResponse(user, nil)
	return c.JSON(res)
}

func (u *UserService) List(c fiber.Ctx) error {
	page := fiber.Query[int](c, "page", 1)
	pagesize := fiber.Query[int](c, "pagesize", 10)

	var datas []dao.User

	dao.Database.Limit(pagesize).Offset((page - 1) * pagesize).Find(&datas)

	res := util.FmtResponse(datas, nil)
	return c.JSON(res)
}

func (u *UserService) Get(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")

	var data dao.User

	dao.Database.First(&data, id)

	res := util.FmtResponse(data, nil)
	return c.JSON(res)
}

func (u *UserService) Create(c fiber.Ctx) error {
	data := new(dao.User)

	if err := c.Bind().Body(data); err != nil {
		return err
	}

	passwd := data.Password
	if passwd == "" {
		return errors.New("password is empty")
	}

	hash := md5.Sum([]byte(passwd))
	hashString := hex.EncodeToString(hash[:])
	data.Password = hashString

	dao.Database.Create(&data)

	res := util.FmtResponse(data, nil)
	return c.JSON(res)
}

func (u *UserService) Update(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")

	var data dao.User

	if err := c.Bind().Body(&data); err != nil {
		return err
	}

	var m dao.User
	dao.Database.First(&m, id)
	dao.Database.Model(&m).Updates(data)

	res := util.FmtResponse(m, nil)

	return c.JSON(res)
}

func (u *UserService) Delete(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")

	dao.Database.Delete(&dao.User{}, id)

	return c.JSON(map[string]int{"code": 200})
}
