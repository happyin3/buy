package auth

import "github.com/gofiber/fiber/v3"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) login(c fiber.Ctx) error {
	username := c.Query("username")
	password := c.Query("password")

	password = utils.Md5String(password)

}
