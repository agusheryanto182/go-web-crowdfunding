package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandlerImpl struct {
	service transaction.TransactionServiceInterface
}

// GetCampaignTransactions implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetCampaignTransactions(c *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateTransaction implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) CreateTransaction(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetCampaignTransactions implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetAllCampaignTransactions(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role != entity.RoleAdmin {
		return response.SendStatusForbidden(c, "forbidden : you do not have permission to access this page")
	}

	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}

	if transactions == nil {
		return response.SendStatusNotFound(c, "transaction is not found")
	}

	return response.SendStatusOkWithDataResponse(c, "success", dto.FormatCampaignTransactions(transactions))
}

// GetNotification implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetNotification(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetUserTransactions implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetUserTransactions(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewTransactionHandler(service transaction.TransactionServiceInterface) transaction.TransactionHandlerInterface {
	return &TransactionHandlerImpl{
		service: service,
	}
}
