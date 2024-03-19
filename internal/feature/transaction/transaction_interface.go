package transaction

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/dto"
	"github.com/gofiber/fiber/v2"
)

type TransactionRepositoryInterface interface {
	GetByCampaignID(page, perPage, campaignID int) ([]*entity.TransactionModels, error)
	GetByUserID(page, perPage, userID int) ([]*entity.TransactionModels, error)
	GetByID(ID int) (*entity.TransactionModels, error)
	Save(transaction *entity.TransactionModels) (*entity.TransactionModels, error)
	Update(transaction *entity.TransactionModels) (*entity.TransactionModels, error)
	FindAll(page, perPage int) ([]*entity.TransactionModels, error)
	GetTotalTransactionCount() (int64, error)
	GetTotalTransactionCountByCampaign(campaignID int) (int64, error)
	GetTotalTransactionCountByUser(userID int) (int64, error)
}

type TransactionServiceInterface interface {
	GetTransactionByCampaignID(page, perPage int, campaignID int) ([]*entity.TransactionModels, int64, error)
	GetTransactionByUserID(page, perPage int, userID int) ([]*entity.TransactionModels, int64, error)
	GetTransactionByID(ID int) (*entity.TransactionModels, error)
	CreateTransaction(payload *dto.CreateTransactionInput) (*entity.TransactionModels, error)
	GetAllTransactions(page, perPage int) ([]*entity.TransactionModels, int64, error)
	ProcessTransactions(input *dto.TransactionNotificationInput) error
}

type TransactionHandlerInterface interface {
	GetAllCampaignTransactions(c *fiber.Ctx) error
	GetCampaignTransactions(c *fiber.Ctx) error
	GetUserTransactions(c *fiber.Ctx) error
	CreateTransaction(c *fiber.Ctx) error
	GetNotification(c *fiber.Ctx) error
}
