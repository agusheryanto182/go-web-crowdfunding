package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
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
	panic("unimplemented")
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
