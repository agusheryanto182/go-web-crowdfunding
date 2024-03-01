package repository

import (
	"errors"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/dto"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) user.UserRepositoryInterface {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (r *UserRepositoryImpl) CreateUser(user *entity.UserModels) (*entity.UserModels, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(userID int, user *dto.UpdateUserRequest) (*entity.UserModels, error) {
	var result *entity.UserModels
	if err := r.DB.Model(&result).Where("id = ?", userID).Updates(&user).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	return result, nil
}

func (r *UserRepositoryImpl) GetByID(userID int) (*entity.UserModels, error) {
	var user entity.UserModels
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
