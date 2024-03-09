package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	DB *gorm.DB
}

// FindByName implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindByName(name string) (*entity.CampaignModels, error) {
	campaign := &entity.CampaignModels{}
	if err := r.DB.Model(&campaign).Where("name = ?", name).First(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

// CreateImage implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) CreateImage(image *entity.CampaignImageModels) (*entity.CampaignImageModels, error) {
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
func (r *CampaignRepositoryImpl) Save(campaign *entity.CampaignModels) (*entity.CampaignModels, error) {
	if err := r.DB.Create(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

// Update implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) Update(campaign *entity.CampaignModels) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

func NewCampaignRepository(DB *gorm.DB) campaign.CampaignRepositoryInterface {
	return &CampaignRepositoryImpl{
		DB: DB,
	}
}
