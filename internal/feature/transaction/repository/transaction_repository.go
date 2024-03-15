package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

// FindAll implements transaction.TransactionRepositoryInterface.
func (t *TransactionRepositoryImpl) FindAll() ([]*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetByCampaignID implements transaction.TransactionRepositoryInterface.
func (t *TransactionRepositoryImpl) GetByCampaignID(campaignID int) ([]*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetByID implements transaction.TransactionRepositoryInterface.
func (t *TransactionRepositoryImpl) GetByID(ID int) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetByUserID implements transaction.TransactionRepositoryInterface.
func (t *TransactionRepositoryImpl) GetByUserID(userID int) ([]*entity.TransactionModels, error) {
	panic("unimplemented")
}

// Save implements transaction.TransactionRepositoryInterface.
func (t *TransactionRepositoryImpl) Save(transaction *entity.TransactionModels) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

// Update implements transaction.TransactionRepositoryInterface.
func (t *TransactionRepositoryImpl) Update(transaction *entity.TransactionModels) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

func NewTransactionRepository(DB *gorm.DB) transaction.TransactionRepositoryInterface {
	return &TransactionRepositoryImpl{
		DB: DB,
	}
}
