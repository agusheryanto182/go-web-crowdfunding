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

func (s *UserServiceImpl) GetByID(userID int) (*entity.UserModels, error) {
	result, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return result, nil
}

func (s *UserServiceImpl) IsAvailableEmail(email string) (bool, error) {
	_, err := s.userRepo.IsAvailableEmail(email)
	if err != nil {
		return false, errors.New("email is not found")
	}

	return true, nil
}
