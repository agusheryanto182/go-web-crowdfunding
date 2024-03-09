package service

import (
	"errors"
	"math"
	"strconv"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/dto"
)

type CampaignServiceImpl struct {
	repo campaign.CampaignRepositoryInterface
}

// FindByNameWithPagination implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) FindByNameWithPagination(page int, perPage int, name string) ([]*entity.CampaignModels, int64, error) {
	campaign, err := s.repo.FindByNameWithPagination(page, perPage, name)
	if err != nil {
		return nil, 0, errors.New("failed to get campaign by name")
	}

	totalItems, err := s.repo.GetTotalCampaignCount()
	if err != nil {
		return nil, 0, errors.New("failed to get total count campaign")
	}

	return campaign, totalItems, nil

}

// CalculatePaginationValues implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
	if page <= 0 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	if page > totalPages {
		page = totalPages
	}

	return page, totalPages
}

// GetNextPage implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) GetNextPage(currentPage int, totalPages int) int {
	if currentPage < totalPages {
		return currentPage + 1
	}
	return totalPages
}

// GetPrevPage implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}

// CreateImage implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) CreateImage(payload *dto.CreateRequestCampaignImage) (*entity.CampaignImageModels, error) {
	panic("unimplemented")
}

// GetAll implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) GetAll(page, perPage int) ([]*entity.CampaignModels, int64, error) {
	campaign, err := s.repo.FindAll(page, perPage)
	if err != nil {
		return nil, 0, errors.New("failed to get all campaign")
	}

	totalItems, err := s.repo.GetTotalCampaignCount()
	if err != nil {
		return nil, 0, errors.New("failed to get count campaign")
	}

	return campaign, totalItems, nil
}

// GetByID implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) GetByID(ID int) (*entity.CampaignModels, error) {
	result, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, errors.New("campaign with ID " + strconv.Itoa(ID) + " is not found")
	}

	return result, nil
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
