package service

import (
	"errors"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/dto"
	utils "github.com/agusheryanto182/go-web-crowdfunding/utils/hash"
)

type UserServiceImpl struct {
	userRepo user.UserRepositoryInterface
	hash     utils.HashInterface
}

func NewUserService(userRepo user.UserRepositoryInterface, hash utils.HashInterface) user.UserServiceInterface {
	return &UserServiceImpl{
		userRepo: userRepo,
		hash:     hash,
	}
}

func (s *UserServiceImpl) CreateUser(payload *dto.RegisterUserRequest) (*entity.UserModels, error) {
	user := entity.UserModels{
		Name:       payload.Name,
		Occupation: payload.Occupation,
		Email:      payload.Email,
	}

	passwordHash, err := s.hash.GenerateHash(payload.Password)
	if err != nil {
		return &entity.UserModels{}, errors.New("failed to generate hash")
	}

	user.Password = passwordHash

	return &user, nil
}

func (s *UserServiceImpl) UpdateUser(userID int, payload *dto.UpdateUserRequest) (*entity.UserModels, error) {
	if _, err := s.userRepo.GetByID(userID); err != nil {
		return nil, errors.New("user not found")
	}

	result, err := s.userRepo.UpdateUser(userID, payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}
