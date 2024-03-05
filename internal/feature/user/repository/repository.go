package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
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

func (r *UserRepositoryImpl) UpdateUser(user *entity.UserModels) (*entity.UserModels, error) {
	err := r.DB.Model(&user).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetByID(userID int) (*entity.UserModels, error) {
	user := &entity.UserModels{}
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindUserByEmail(email string) (*entity.UserModels, error) {
	user := &entity.UserModels{}
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
