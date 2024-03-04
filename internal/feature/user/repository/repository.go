package repository

import (
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

func (r *UserRepositoryImpl) UpdateUser(userID int, user *dto.UpdateUserRequest) (*entity.UserModels, error) {
	var result *entity.UserModels
	if err := r.DB.Model(&result).Where("id = ?", userID).Updates(&user).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepositoryImpl) GetByID(userID int) (*entity.UserModels, error) {
	var user *entity.UserModels
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) IsAvailableEmail(email string) (*entity.UserModels, error) {
	var user *entity.UserModels
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
