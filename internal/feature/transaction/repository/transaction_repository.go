package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// GetTotalTransactionCountByCampaign implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetTotalTransactionCountByCampaign(campaignID int) (int64, error) {
	var totalItems int64
	if err := r.db.Model(&entity.TransactionModels{}).Where("campaign_id = ?", campaignID).Count(&totalItems).Error; err != nil {
		return 0, err
	}
	return totalItems, nil
}

// GetTotalTransactionCountByUser implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetTotalTransactionCountByUser(userID int) (int64, error) {
	var totalItems int64
	if err := r.db.Model(&entity.TransactionModels{}).Where("user_id = ?", userID).Count(&totalItems).Error; err != nil {
		return 0, err
	}
	return totalItems, nil
}

// GetTotalTransactionCount implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetTotalTransactionCount() (int64, error) {
	var totalItems int64
	if err := r.db.Model(&entity.TransactionModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}
	return totalItems, nil
}

// FindAll implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) FindAll(page, perPage int) ([]*entity.TransactionModels, error) {
	var transactions []*entity.TransactionModels
	offset := (page - 1) * perPage
	query := r.db.Offset(offset).Limit(perPage)

	if err := query.Preload("Campaigns").Preload("Users").Order("id desc").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetByCampaignID implements transaction.TransactionRepositoryInterface.
func (r *TransactionRepositoryImpl) GetByCampaignID(page, perPage, campaignID int) ([]*entity.TransactionModels, error) {
	var transaction []*entity.TransactionModels
	offset := (page - 1) * perPage
	query := r.db.Offset(offset).Limit(perPage)

	if err := query.Preload("Campaigns").Preload("Users").Where("campaign_id = ?", campaignID).Find(&transaction).Error; err != nil {
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
func (r *TransactionRepositoryImpl) GetByUserID(page, perPage, userID int) ([]*entity.TransactionModels, error) {
	var transaction []*entity.TransactionModels
	offset := (page - 1) * perPage
	query := r.db.Offset(offset).Limit(perPage)

	if err := query.Preload("Campaigns").Where("user_id = ?", userID).Find(&transaction).Error; err != nil {
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
