package service

import (
	"errors"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	utils "github.com/agusheryanto182/go-web-crowdfunding/utils/hash"
)

type AuthServiceImpl struct {
	authRepo    auth.AuthRepositoryInterface
	userService user.UserServiceInterface
	hash        utils.HashInterface
}

func NewAuthService(authRepo auth.AuthRepositoryInterface, userService user.UserServiceInterface, hash utils.HashInterface) auth.AuthServiceInterface {
	return &AuthServiceImpl{
		authRepo:    authRepo,
		userService: userService,
		hash:        hash,
	}
}

func (s *AuthServiceImpl) SignUp(payload *dto.RegisterUserRequest) (*entity.UserModels, error) {
	isExistEmail, _ := s.userService.IsAvailableEmail(payload.Email)
	if isExistEmail {
		return nil, errors.New("email is already exist")
	}

	user := &entity.UserModels{
		Name:       payload.Name,
		Occupation: payload.Occupation,
		Email:      payload.Email,
		Role:       "user",
	}

	passwordHash, err := s.hash.GenerateHash(payload.Password)
	if err != nil {
		return &entity.UserModels{}, errors.New("failed to generate hash")
	}

	user.Password = passwordHash

	result, err := s.authRepo.SignUp(user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return result, nil
}
