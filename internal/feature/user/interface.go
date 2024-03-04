package user

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/dto"
	"github.com/gofiber/fiber/v2"
)

type UserRepositoryInterface interface {
	UpdateUser(userID int, user *dto.UpdateUserRequest) (*entity.UserModels, error)
	GetByID(userID int) (*entity.UserModels, error)
	IsAvailableEmail(email string) (*entity.UserModels, error)
}

type UserServiceInterface interface {
	UpdateUser(userID int, payload *dto.UpdateUserRequest) (*entity.UserModels, error)
	GetByID(userID int) (*entity.UserModels, error)
	IsAvailableEmail(email string) (bool, error)
}

type UserHandlerInterface interface {
	Update(c *fiber.Ctx) error
}
