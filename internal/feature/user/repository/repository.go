package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// DeleteUser implements user.UserRepositoryInterface.
func (r *UserRepositoryImpl) DeleteUser(ID int) error {
	if err := r.DB.Where("id = ?", ID).Delete(&entity.UserModels{}).Error; err != nil {
		return err
	}
	return nil
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

func (r *UserRepositoryImpl) FindAllUser(page, perPage int) ([]*entity.UserModels, error) {
	var user []*entity.UserModels
	offset := (page - 1) * perPage
	if err := r.DB.Offset(offset).Limit(perPage).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindUserByName(page, perPage int, name string) ([]*entity.UserModels, error) {
	var user []*entity.UserModels
	offset := (page - 1) * perPage
	query := r.DB.Offset(offset).Limit(perPage)
	if err := query.Where("name LIKE ?", "%"+name+"%").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetTotalUserCount() (int64, error) {
	var totalItems int64

	if err := r.DB.Model(&entity.UserModels{}).Where("is_verified = ?", true).Count(&totalItems).Error; err != nil {
		return 0, err
	}
	return totalItems, nil
}

func NewUserRepository(DB *gorm.DB) user.UserRepositoryInterface {
	return &UserRepositoryImpl{
		DB: DB,
	}
}
