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
	panic("unimplemented")
}

// GetByID implements transaction.TransactionRepositoryInterface.
func (t *TransactionRepositoryImpl) GetByID(ID int) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetByUserID implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetByUserID(userID int) ([]*entity.TransactionModels, error) {
	panic("unimplemented")
}

// Save implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) Save(transaction *entity.TransactionModels) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

// Update implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) Update(transaction *entity.TransactionModels) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepositoryInterface {
	return &TransactionRepositoryImpl{
		db: db,
	}
}
