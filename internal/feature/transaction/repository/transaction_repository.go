package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// FindAll implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) FindAll() ([]*entity.TransactionModels, error) {
	var transactions []*entity.TransactionModels
	if err := r.db.Preload("Campaigns").Preload("Users").Order("id desc").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByCampaignID implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetByCampaignID(campaignID int) ([]*entity.TransactionModels, error) {
	var transaction []*entity.TransactionModels

	if err := r.db.Preload("Campaigns").Preload("Users").Where("campaign_id = ?", campaignID).Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

// GetByID implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetByID(ID int) (*entity.TransactionModels, error) {
	transaction := &entity.TransactionModels{}
	if err := r.db.Where("id = ?", ID).First(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// GetByUserID implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetByUserID(userID int) ([]*entity.TransactionModels, error) {
	var transaction []*entity.TransactionModels

	if err := r.db.Preload("Campaigns").Where("user_id = ?", userID).Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// Save implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) Save(transaction *entity.TransactionModels) (*entity.TransactionModels, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// Update implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) Update(transaction *entity.TransactionModels) (*entity.TransactionModels, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepositoryInterface {
	return &TransactionRepositoryImpl{
		db: db,
	}
}
