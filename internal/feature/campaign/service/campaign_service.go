package service

import (
	"errors"
	"math"
	"strconv"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
)

type CampaignServiceImpl struct {
	repo        campaign.CampaignRepositoryInterface
	userService user.UserServiceInterface
}

// DeleteCampaign implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) DeleteCampaign(ID int) error {
	if err := s.repo.DeleteCampaign(ID); err != nil {
		return errors.New("failed to delete campaign : " + err.Error())
	}
	return nil
}

// DeleteImageCampaign implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) DeleteImageCampaign(campaignID, imageID int) error {
	if err := s.repo.DeleteImageCampaign(campaignID, imageID); err != nil {
		return errors.New("failed to delete image campaign")
	}
	return nil
}

// FindImageByID implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) FindImageByID(ID int) (*entity.CampaignImageModels, error) {
	image, err := s.repo.FindImageByID(ID)
	if err != nil {
		return nil, errors.New("image with ID = " + strconv.Itoa(ID) + " is not found")
	}
	return image, nil
}

// SetPrimaryImage implements campaign.CampaignServiceInterface.
func (s *CampaignServiceImpl) SetPrimaryImage(payload *dto.SetPrimaryImageRequest) (*entity.CampaignImageModels, error) {
	image := &entity.CampaignImageModels{
		ID:         payload.ID,
		CampaignID: payload.CampaignID,
		IsPrimary:  1,
	}

	campaignImages, err := s.repo.FindAllImagesCampaign(payload.CampaignID)
	if err != nil {
		return nil, errors.New("failed get all images by campaign id")
	}

	for _, images := range campaignImages {
		if images.IsPrimary == 1 {
			images.IsPrimary = 0
			_, err := s.repo.SetPrimaryImage(images)
			if err != nil {
				return nil, errors.New("failed set image : " + err.Error())
			}
		}
	}

	result, err := s.repo.SetPrimaryImage(image)
	if err != nil {
		return nil, errors.New("failed set image to primary : " + err.Error())
	}
	return result, nil
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
	campaignImage := &entity.CampaignImageModels{
		CampaignID: payload.CampaignID,
		FileName:   payload.FileName,
		IsPrimary:  0,
	}

	result, err := s.repo.CreateImage(campaignImage)
	if err != nil {
		return nil, errors.New("failed to create image")
	}
	return result, nil
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
func (s *CampaignServiceImpl) GetByUserID(page, perPage, UserID int, name string) ([]*entity.CampaignModels, int64, error) {
	if _, err := s.userService.GetByID(UserID); err != nil {
		return nil, 0, err
	}

	campaigns, err := s.repo.FindByUserID(page, perPage, UserID, name)
	if err != nil {
		return nil, 0, errors.New("failed to get campaign : " + err.Error())
	}

	if len(campaigns) == 0 {
		return nil, 0, errors.New("failed to get campaign with userID : " + strconv.Itoa(UserID))
	}

	totalItems, err := s.repo.GetTotalCampaignCount()
	if err != nil {
		return nil, 0, errors.New("failed to get campaign count")
	}

	return campaigns, totalItems, nil
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
func (s *CampaignServiceImpl) Update(userID int, payload *dto.UpdateRequestCampaign) (*entity.CampaignModels, error) {
	campaign, _ := s.GetByID(payload.ID)
	if campaign == nil {
		return nil, errors.New("campaign with ID : " + strconv.Itoa(payload.ID) + " is not found")
	}

	if campaign.UserID != userID {
		return nil, errors.New("invalid credentials")
	}

	campaign.Name = payload.Name
	campaign.ShortDescription = payload.ShortDescription
	campaign.Description = payload.Description
	campaign.Perks = payload.Perks
	campaign.GoalAmount = payload.GoalAmount

	result, err := s.repo.Update(campaign)
	if err != nil {
		return nil, errors.New("failed to update campaign : " + err.Error())
	}

	return result, nil
}

func NewCampaignService(repo campaign.CampaignRepositoryInterface, userService user.UserServiceInterface) campaign.CampaignServiceInterface {
	return &CampaignServiceImpl{
		repo:        repo,
		userService: userService,
	}
}
