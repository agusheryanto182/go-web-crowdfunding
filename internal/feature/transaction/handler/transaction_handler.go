package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandlerImpl struct {
	service transaction.TransactionServiceInterface
}

// CreateTransaction implements transaction.TransactionHandlerInterface.
func (t *TransactionHandlerImpl) CreateTransaction(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetCampaignTransactions implements transaction.TransactionHandlerInterface.
func (t *TransactionHandlerImpl) GetCampaignTransactions(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetNotification implements transaction.TransactionHandlerInterface.
func (t *TransactionHandlerImpl) GetNotification(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetUserTransactions implements transaction.TransactionHandlerInterface.
func (t *TransactionHandlerImpl) GetUserTransactions(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewTransactionHandler(service transaction.TransactionServiceInterface) transaction.TransactionHandlerInterface {
	return &TransactionHandlerImpl{
		service: service,
	}
}
