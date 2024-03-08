package user

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/dto"
	"github.com/gofiber/fiber/v2"
)

type UserRepositoryInterface interface {
	UpdateUser(user *entity.UserModels) (*entity.UserModels, error)
	GetByID(userID int) (*entity.UserModels, error)
	FindUserByEmail(email string) (*entity.UserModels, error)
	UploadAvatar(userID int, avatarPath string) (*entity.UserModels, error)
	GetAllUser() ([]*entity.UserModels, error)
}

type UserServiceInterface interface {
	UpdateUser(userID int, payload *dto.UpdateUserRequest) (*entity.UserModels, error)
	GetByID(userID int) (*entity.UserModels, error)
	GetUserByEmail(email string) (*entity.UserModels, error)
	UploadAvatar(userID int, avatar *dto.UpdateAvatarRequest) (*entity.UserModels, error)
	GetAllUser() ([]*entity.UserModels, error)
}

type UserHandlerInterface interface {
	Update(c *fiber.Ctx) error
	UploadAvatar(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
}
