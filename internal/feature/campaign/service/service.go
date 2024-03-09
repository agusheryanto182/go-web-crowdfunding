package service

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/dto"
)

type CampaignServiceImpl struct {
	repo campaign.CampaignRepositoryInterface
}

// CreateImage implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) CreateImage(payload *dto.CreateRequestCampaignImage) (*entity.CampaignImageModels, error) {
	panic("unimplemented")
}

// GetAll implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) GetAll() (*entity.CampaignModels, error) {
	panic("unimplemented")
}

// GetByID implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) GetByID(ID int) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

// GetByUserID implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) GetByUserID(UserID int) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

// Save implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) Save(payload *dto.CreateRequestCampaign) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

// Update implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) Update(payload *dto.UpdateRequestCampaign) (*entity.CampaignModels, error) {
	panic("unimplemented")
}

func NewCampaignService(repo campaign.CampaignRepositoryInterface) campaign.CampaignServiceInterface {
	return &CampaignServiceImpl{
		repo: repo,
	}
}
