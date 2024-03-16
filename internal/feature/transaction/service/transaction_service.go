package service

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/dto"
)

type TransactionServiceImpl struct {
	repo transaction.TransactionRepositoryInterface
}

// CreateTransaction implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) CreateTransaction(payload *dto.CreateTransactionInput) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetAllTransactions implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetAllTransactions() ([]*entity.TransactionModels, error) {
	transactions, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionByCampaignID implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetTransactionByCampaignID(campaignID int) ([]*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetTransactionByID implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetTransactionByID(ID int) (*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetTransactionByUserID implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetTransactionByUserID(userID int) ([]*entity.TransactionModels, error) {
	panic("unimplemented")
}

func NewTransactionService(repo transaction.TransactionRepositoryInterface) transaction.TransactionServiceInterface {
	return &TransactionServiceImpl{
		repo: repo,
	}
}
