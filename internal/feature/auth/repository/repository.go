package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) auth.AuthRepositoryInterface {
	return &AuthRepositoryImpl{
		DB: DB,
	}
}

func (r *AuthRepositoryImpl) SignUp(user *entity.UserModels) (*entity.UserModels, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AuthRepositoryImpl) SaveOTP(OTP *entity.OTPModels) (*entity.OTPModels, error) {
	if err := r.DB.Create(&OTP).Error; err != nil {
		return nil, err
	}
	return OTP, nil
}
