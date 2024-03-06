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
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.Name = payload.Name
	user.Email = payload.Email
	user.Occupation = payload.Occupation
	user.Password = payload.Password

	result, err := s.userRepo.UpdateUser(user)
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

func (s *UserServiceImpl) GetUserByEmail(email string) (*entity.UserModels, error) {
	result, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("email is not found")
	}

	return result, nil
}

func (s *UserServiceImpl) UploadAvatar(userID int, avatar *dto.UpdateAvatarRequest) (*entity.UserModels, error) {
	result, err := s.userRepo.UploadAvatar(userID, avatar.Avatar)
	if err != nil {
		return nil, errors.New("failed to upload avatar : " + err.Error())
	}

	return result, nil
}
