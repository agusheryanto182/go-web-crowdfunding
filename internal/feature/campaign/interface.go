package campaign

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/dto"
	"github.com/gofiber/fiber/v2"
)

type CampaignRepositoryInterface interface {
	FindAll(page, perPage int) ([]*entity.CampaignModels, error)
	FindByUserID(page, perPage, userID int, name string) ([]*entity.CampaignModels, error)
	FindByID(ID int) (*entity.CampaignModels, error)
	Save(campaign *entity.CampaignModels) (*entity.CampaignModels, error)
	Update(campaign *entity.CampaignModels) (*entity.CampaignModels, error)
	CreateImage(image *entity.CampaignImageModels) (*entity.CampaignImageModels, error)
	FindByName(name string) (*entity.CampaignModels, error)
	GetTotalCampaignCount() (int64, error)
	FindByNameWithPagination(page, perPage int, name string) ([]*entity.CampaignModels, error)
}

type CampaignServiceInterface interface {
	GetAll(page, perPage int) ([]*entity.CampaignModels, int64, error)
	GetByUserID(page, perPage, UserID int, name string) ([]*entity.CampaignModels, int64, error)
	GetByID(ID int) (*entity.CampaignModels, error)
	Save(payload *dto.CreateRequestCampaign) (*entity.CampaignModels, error)
	Update(payload *dto.UpdateRequestCampaign) (*entity.CampaignModels, error)
	CreateImage(payload *dto.CreateRequestCampaignImage) (*entity.CampaignImageModels, error)
	CalculatePaginationValues(page int, totalItems int, perPage int) (int, int)
	GetNextPage(currentPage int, totalPages int) int
	GetPrevPage(currentPage int) int
	FindByNameWithPagination(page, perPage int, name string) ([]*entity.CampaignModels, int64, error)
}

type CampaignHandlerInterface interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	CreateImage(c *fiber.Ctx) error
}
