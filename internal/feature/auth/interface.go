package auth

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthRepositoryInterface interface {
	SignUp(user *entity.UserModels) (*entity.UserModels, error)
	SaveOTP(OTP *entity.OTPModels) (*entity.OTPModels, error)
}

type AuthServiceInterface interface {
	SignUp(user *dto.RegisterUserRequest) (*entity.UserModels, error)
}

type AuthHandlerInterface interface {
	SignUp(c *fiber.Ctx) error
}
