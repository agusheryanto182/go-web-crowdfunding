package transaction

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/dto"
	"github.com/gofiber/fiber/v2"
)

type TransactionRepositoryInterface interface {
	GetByCampaignID(campaignID int) ([]*entity.TransactionModels, error)
	GetByUserID(userID int) ([]*entity.TransactionModels, error)
	GetByID(ID int) (*entity.TransactionModels, error)
	Save(transaction *entity.TransactionModels) (*entity.TransactionModels, error)
	Update(transaction *entity.TransactionModels) (*entity.TransactionModels, error)
	FindAll() ([]*entity.TransactionModels, error)
}

type TransactionServiceInterface interface {
	GetTransactionByCampaignID(campaignID int) ([]*entity.TransactionModels, error)
	GetTransactionByUserID(userID int) ([]*entity.TransactionModels, error)
	GetTransactionByID(ID int) (*entity.TransactionModels, error)
	CreateTransaction(payload *dto.CreateTransactionInput) (*entity.TransactionModels, error)
	GetAllTransactions() ([]*entity.TransactionModels, error)
	ProcessTransactions(input *dto.TransactionNotificationInput) error
}

type TransactionHandlerInterface interface {
	GetAllCampaignTransactions(c *fiber.Ctx) error
	GetCampaignTransactions(c *fiber.Ctx) error
	GetUserTransactions(c *fiber.Ctx) error
	CreateTransaction(c *fiber.Ctx) error
	GetNotification(c *fiber.Ctx) error
}
