package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	DB *gorm.DB
}

// CreateImage implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) CreateImage(input *entity.CampaignImageModels) (*entity.CampaignImageModels, error) {
	panic("unimplemented")
}

// FindAll implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindAll() (*entity.CampaignModels, error) {
	panic("unimplemented")
}

// FindByID implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindByID(ID int) (*entity.UserModels, error) {
	panic("unimplemented")
}

// FindByUserID implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindByUserID(userID int) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

// Save implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) Save(input *entity.CampaignModels) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

// Update implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) Update(input *entity.CampaignModels) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

func NewCampaignRepository(DB *gorm.DB) campaign.CampaignRepositoryInterface {
	return &CampaignRepositoryImpl{
		DB: DB,
	}
}
