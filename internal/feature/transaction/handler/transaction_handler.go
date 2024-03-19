package handler

import (
	"strconv"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandlerImpl struct {
	service         transaction.TransactionServiceInterface
	campaignService campaign.CampaignServiceInterface
}

// GetCampaignTransactions implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetCampaignTransactions(c *fiber.Ctx) error {
	campaignID, _ := strconv.Atoi(c.Params("id"))
	page, _ := strconv.Atoi(c.Query("page"))
	perPage := 7

	var transactions []*entity.TransactionModels

	var totalItems int64
	var err error

	campaign, err := h.campaignService.GetByID(campaignID)
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	if campaign == nil {
		return response.SendStatusNotFound(c, "not found : "+err.Error())
	}

	transactions, totalItems, err = h.service.GetTransactionByCampaignID(page, perPage, campaignID)
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	if transactions == nil {
		return response.SendStatusNotFound(c, "not found : "+err.Error())
	}

	currentPage, totalPages := h.campaignService.CalculatePaginationValues(page, int(totalItems), perPage)
	nextPage := h.campaignService.GetNextPage(currentPage, totalPages)
	prevPage := h.campaignService.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.FormatCampaignTransactions(transactions), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success")
}

// CreateTransaction implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) CreateTransaction(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role == entity.RoleAdmin {
		return response.SendStatusForbidden(c, "forbidden : you do not have permission to access this page")
	}

	var payload dto.CreateTransactionInput
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid input")
	}

	if err := validator.ValidateStruct(&payload); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	checkCampaign, err := h.campaignService.GetByID(payload.CampaignID)
	if err != nil {
		return response.SendStatusNotFound(c, err.Error())
	}

	if currentUser.ID == checkCampaign.UserID {
		return response.SendStatusForbidden(c, "forbidden : you cannot pay your self :)")
	}

	payload.User = *currentUser

	result, err := h.service.CreateTransaction(&payload)
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", dto.FormatTransaction(result))

}

// GetCampaignTransactions implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetAllCampaignTransactions(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role != entity.RoleAdmin {
		return response.SendStatusForbidden(c, "forbidden : you do not have permission to access this page")
	}

	page, _ := strconv.Atoi(c.Query("page"))
	perPage := 7

	var transactions []*entity.TransactionModels

	var totalItems int64
	var err error

	transactions, totalItems, err = h.service.GetAllTransactions(page, perPage)
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}

	if transactions == nil {
		return response.SendStatusNotFound(c, "transaction is not found")
	}

	currentPage, totalPages := h.campaignService.CalculatePaginationValues(page, int(totalItems), perPage)
	nextPage := h.campaignService.GetNextPage(currentPage, totalPages)
	prevPage := h.campaignService.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.FormatCampaignTransactions(transactions), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success")

}

// GetNotification implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetNotification(c *fiber.Ctx) error {
	var input *dto.TransactionNotificationInput

	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	if err := h.service.ProcessTransactions(input); err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", input)
}

// GetUserTransactions implements transaction.TransactionHandlerInterface.
func (h *TransactionHandlerImpl) GetUserTransactions(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role == entity.RoleAdmin {
		return response.SendStatusForbidden(c, "forbidden : you do not have access for this route")
	}

	page, _ := strconv.Atoi(c.Query("page"))
	perPage := 7

	var transactions []*entity.TransactionModels

	var totalItems int64
	var err error

	transactions, totalItems, err = h.service.GetTransactionByUserID(page, perPage, currentUser.ID)
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	currentPage, totalPages := h.campaignService.CalculatePaginationValues(page, int(totalItems), perPage)
	nextPage := h.campaignService.GetNextPage(currentPage, totalPages)
	prevPage := h.campaignService.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.FormatCampaignTransactions(transactions), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success")

}

func NewTransactionHandler(service transaction.TransactionServiceInterface, campaignService campaign.CampaignServiceInterface) transaction.TransactionHandlerInterface {
	return &TransactionHandlerImpl{
		service:         service,
		campaignService: campaignService,
	}
}
