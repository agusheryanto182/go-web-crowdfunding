package user

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/dto"
	"github.com/gofiber/fiber/v2"
)

type UserRepositoryInterface interface {
	CreateUser(user *entity.UserModels) (*entity.UserModels, error)
	UpdateUser(userID int, user *dto.UpdateUserRequest) (*entity.UserModels, error)
	GetByID(userID int) (*entity.UserModels, error)
}

type UserServiceInterface interface {
	CreateUser(payload *dto.RegisterUserRequest) (*entity.UserModels, error)
	UpdateUser(userID int, payload *dto.UpdateUserRequest) (*entity.UserModels, error)
}

type UserHandlerInterface interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
