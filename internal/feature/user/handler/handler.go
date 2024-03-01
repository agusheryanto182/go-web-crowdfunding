package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerImpl struct {
	userService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) user.UserHandlerInterface {
	return &UserHandlerImpl{userService: userService}
}

func (u *UserHandlerImpl) Create(c *fiber.Ctx) error {
	panic("unimplemented")
}

func (u *UserHandlerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}
