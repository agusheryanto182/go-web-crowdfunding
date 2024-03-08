package service

import (
	"errors"
	"math"

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

func (s *UserServiceImpl) GetAllUser(page, perPage int) ([]*entity.UserModels, int64, error) {
	result, err := s.userRepo.FindAllUser(page, perPage)
	if err != nil {
		return nil, 0, errors.New("failed to get all user")
	}

	totalItems, err := s.userRepo.GetTotalUserCount()
	if err != nil {
		return nil, 0, errors.New("failed to get total user")
	}

	return result, totalItems, nil
}

func (s *UserServiceImpl) GetUserByName(page, perPage int, name string) ([]*entity.UserModels, int64, error) {
	user, err := s.userRepo.FindUserByName(page, perPage, name)
	if err != nil {
		return nil, 0, errors.New("failed to get user by name")
	}

	totalItems, err := s.userRepo.GetTotalUserCount()
	if err != nil {
		return nil, 0, errors.New("failed to get total user")
	}

	return user, totalItems, nil
}

func (s *UserServiceImpl) CalculatePaginationValues(page, totalItems, perPage int) (int, int) {
	if page <= 0 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if page > totalPages {
		page = totalPages
	}

	return page, totalPages
}

func (s *UserServiceImpl) GetNextPage(currentPage int, totalPages int) int {
	if currentPage < totalPages {
		return currentPage + 1
	}

	return totalPages
}

func (s *UserServiceImpl) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}

	return 1
}
