package handler

import (
	"strconv"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type CampaignHandlerImpl struct {
	service campaign.CampaignServiceInterface
}

// CreateImage implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) CreateImage(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage := 5

	var campaign []*entity.CampaignModels
	var totalItems int64
	var err error
	search := c.Query("search")

	if search != "" {
		campaign, totalItems, err = h.service.FindByNameWithPagination(page, perPage, search)
	} else {
		campaign, totalItems, err = h.service.GetAll(page, perPage)
	}

	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	currentPage, totalPages := h.service.CalculatePaginationValues(page, int(totalItems), perPage)
	nextPage := h.service.GetNextPage(currentPage, totalPages)
	prevPage := h.service.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.FormatCampaignsResponse(campaign), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success")

}

// GetByID implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) GetByID(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}
	campaign, err := h.service.GetByID(ID)
	if err != nil {
		return response.SendStatusNotFound(c, "error : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", *dto.FormatSaveCampaignResponse(campaign))
}

// GetByUserID implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) GetByUserID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Save implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) Save(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser == nil {
		return response.SendStatusInternalServerError(c, "something error with middleware")
	}

	campaign := &dto.CreateRequestCampaign{}
	if err := c.BodyParser(campaign); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	if err := validator.ValidateStruct(campaign); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	campaign.UserID = currentUser.ID

	result, err := h.service.Save(campaign)
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", *dto.FormatSaveCampaignResponse(result))

}

// Update implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewCampaignHandler(service campaign.CampaignServiceInterface) campaign.CampaignHandlerInterface {
	return &CampaignHandlerImpl{
		service: service,
	}
}
