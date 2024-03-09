package service

import (
	"errors"

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
	checkName, _ := s.repo.FindByName(payload.Name)
	if checkName != nil {
		return nil, errors.New("name is already exist")
	}

	campaign := &entity.CampaignModels{
		UserID:           payload.UserID,
		Name:             payload.Name,
		ShortDescription: payload.ShortDescription,
		Description:      payload.Description,
		Perks:            payload.Perks,
		GoalAmount:       payload.GoalAmount,
	}

	result, err := s.repo.Save(campaign)
	if err != nil {
		return nil, errors.New("failed to save campaign")
	}

	return result, nil
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
