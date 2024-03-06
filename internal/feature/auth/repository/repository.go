package repository

import (
	"time"

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

func (r *AuthRepositoryImpl) FindValidOTP(userID int, OTP string) (*entity.OTPModels, error) {
	var validOTP *entity.OTPModels
	if err := r.DB.Where("user_id = ? AND otp = ? AND expired_otp > ?", userID, OTP, time.Now().Unix()).First(&validOTP).Error; err != nil {
		return validOTP, err
	}
	return validOTP, nil
}

func (r *AuthRepositoryImpl) DeleteOTP(OTP *entity.OTPModels) error {
	if err := r.DB.Where("id = ?", OTP.ID).Delete(&OTP).Error; err != nil {
		return err
	}
	return nil
}
