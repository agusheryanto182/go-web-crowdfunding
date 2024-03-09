package campaign

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/dto"
	"github.com/gofiber/fiber/v2"
)

type CampaignRepositoryInterface interface {
	FindAll() (*entity.CampaignModels, error)
	FindByUserID(userID int) (*entity.CampaignModels, error)
	FindByID(ID int) (*entity.UserModels, error)
	Save(input *entity.CampaignModels) (*entity.CampaignModels, error)
	Update(input *entity.CampaignModels) (*entity.CampaignModels, error)
	CreateImage(input *entity.CampaignImageModels) (*entity.CampaignImageModels, error)
}

type CampaignServiceInterface interface {
	GetAll() (*entity.CampaignModels, error)
	GetByUserID(UserID int) (*entity.CampaignModels, error)
	GetByID(ID int) (*entity.CampaignModels, error)
	Save(payload *dto.CreateRequestCampaign) (*entity.CampaignModels, error)
	Update(payload *dto.UpdateRequestCampaign) (*entity.CampaignModels, error)
	CreateImage(payload *dto.CreateRequestCampaignImage) (*entity.CampaignImageModels, error)
}

type CampaignHandlerInterface interface {
	GetAll(c *fiber.Ctx) error
	GetByUserID(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	CreateImage(c *fiber.Ctx) error
}
