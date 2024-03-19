package service

import (
	"strconv"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/payment"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/dto"
)

type TransactionServiceImpl struct {
	repo               transaction.TransactionRepositoryInterface
	paymentService     payment.PaymentServiceInterface
	campaignRepository campaign.CampaignRepositoryInterface
	campaignService    campaign.CampaignServiceInterface
}

// ProcessTransactions implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) ProcessTransactions(input *dto.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repo.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "canceled"
	}

	updatedTransaction, err := s.repo.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateTransaction implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) CreateTransaction(payload *dto.CreateTransactionInput) (*entity.TransactionModels, error) {
	transaction := &entity.TransactionModels{
		Amount:     payload.Amount,
		UserID:     payload.User.ID,
		CampaignID: payload.CampaignID,
		Status:     "pending",
	}
	newTransaction, err := s.repo.Save(transaction)
	if err != nil {
		return nil, err
	}

	paymentURL, err := s.paymentService.GetPaymentURL(*newTransaction, payload.User)
	if err != nil {
		return nil, err
	}

	newTransaction.PaymentURL = paymentURL

	newResult, err := s.repo.Update(newTransaction)
	if err != nil {
		return nil, err
	}

	return newResult, nil
}

// GetAllTransactions implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetAllTransactions(page, perPage int) ([]*entity.TransactionModels, int64, error) {
	transactions, err := s.repo.FindAll(page, perPage)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalTransactionCount()
	if err != nil {
		return nil, 0, err
	}
	return transactions, totalItems, nil
}

// GetTransactionByCampaignID implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetTransactionByCampaignID(page, perPage, campaignID int) ([]*entity.TransactionModels, int64, error) {
	transactions, err := s.repo.GetByCampaignID(page, perPage, campaignID)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalTransactionCountByCampaign(campaignID)
	if err != nil {
		return nil, 0, err
	}
	return transactions, totalItems, nil
}

// GetTransactionByID implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetTransactionByID(ID int) (*entity.TransactionModels, error) {
	transactions, err := s.repo.GetByID(ID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionByUserID implements transaction.TransactionServiceInterface.
func (s *TransactionServiceImpl) GetTransactionByUserID(page, perPage, userID int) ([]*entity.TransactionModels, int64, error) {
	transactions, err := s.repo.GetByUserID(page, perPage, userID)
	if err != nil {
		return nil, 0, err
	}
	totalItems, err := s.repo.GetTotalTransactionCountByUser(userID)
	if err != nil {
		return nil, 0, err
	}
	return transactions, totalItems, nil
}

func NewTransactionService(
	repo transaction.TransactionRepositoryInterface,
	paymentService payment.PaymentServiceInterface,
	campaignRepository campaign.CampaignRepositoryInterface,
	campaignService campaign.CampaignServiceInterface) transaction.TransactionServiceInterface {
	return &TransactionServiceImpl{
		repo:               repo,
		paymentService:     paymentService,
		campaignRepository: campaignRepository,
		campaignService:    campaignService,
	}
}
