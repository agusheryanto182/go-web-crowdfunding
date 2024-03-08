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
	if err := r.DB.Where("id = ? AND is_verified = ?", userID, true).First(&user).Error; err != nil {
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

func (r *UserRepositoryImpl) UploadAvatar(userID int, avatar string) (*entity.UserModels, error) {
	user := &entity.UserModels{}
	if err := r.DB.Model(&user).Where("id = ?", userID).UpdateColumn("avatar", avatar).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetAllUser() ([]*entity.UserModels, error) {
	var user []*entity.UserModels
	if err := r.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
