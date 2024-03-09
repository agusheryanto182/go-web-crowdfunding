package handler

import (
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
	panic("unimplemented")
}

// GetByID implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) GetByID(c *fiber.Ctx) error {
	panic("unimplemented")
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
