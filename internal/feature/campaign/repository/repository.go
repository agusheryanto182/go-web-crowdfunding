package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	DB *gorm.DB
}

// DeleteCampaign implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) DeleteCampaign(ID int) error {
	if err := r.DB.Where("id = ?", ID).Delete(&entity.CampaignModels{}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteImageCampaign implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) DeleteImageCampaign(campaignID, imageID int) error {
	image := &entity.CampaignImageModels{}
	if err := r.DB.Where("id = ?", imageID).First(&image).Error; err != nil {
		return err
	}

	if err := r.DB.Where("id = ? AND campaign_id = ?", imageID, campaignID).Delete(&image).Error; err != nil {
		return err
	}

	return nil
}

// FindImageByID implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindImageByID(ID int) (*entity.CampaignImageModels, error) {
	image := &entity.CampaignImageModels{}
	if err := r.DB.Where("id = ?", ID).First(image).Error; err != nil {
		return nil, err
	}
	return image, nil
}

// FindAllImages implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindAllImagesCampaign(campaignID int) ([]*entity.CampaignImageModels, error) {
	var campaignImages []*entity.CampaignImageModels
	if err := r.DB.Where("campaign_id = ?", campaignID).Find(&campaignImages).Error; err != nil {
		return nil, err
	}
	return campaignImages, nil
}

// SetPrimaryImage implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) SetPrimaryImage(image *entity.CampaignImageModels) (*entity.CampaignImageModels, error) {
	if err := r.DB.Model(&image).Where("id = ? AND campaign_id = ?", image.ID, image.CampaignID).Update("is_primary", image.IsPrimary).Error; err != nil {
		return nil, err
	}
	return image, nil
}

// FindByNameWithPagination implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindByNameWithPagination(page int, perPage int, name string) ([]*entity.CampaignModels, error) {
	campaign := []*entity.CampaignModels{}
	offset := (page - 1) * perPage
	if err := r.DB.Offset(offset).Limit(perPage).Where("name LIKE ?", "%"+name+"%").Find(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

// GetTotalUserCount implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) GetTotalCampaignCount() (int64, error) {
	var totalItems int64
	if err := r.DB.Model(&entity.CampaignModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}
	return totalItems, nil
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
	if err := r.DB.Create(&image).Error; err != nil {
		return nil, err
	}
	return image, nil
}

// FindAll implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindAll(page, perPage int) ([]*entity.CampaignModels, error) {
	var campaign []*entity.CampaignModels
	offset := (page - 1) * perPage
	if err := r.DB.Offset(offset).Limit(perPage).Find(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

// FindByID implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindByID(ID int) (*entity.CampaignModels, error) {
	campaign := &entity.CampaignModels{}
	if err := r.DB.Model(&campaign).Where("id = ?", ID).First(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

// FindByUserID implements campaign.CampaignRepositoryInterface.
func (r *CampaignRepositoryImpl) FindByUserID(page, perPage, userID int, name string) ([]*entity.CampaignModels, error) {
	var campaigns []*entity.CampaignModels
	offset := (page - 1) * perPage
	query := r.DB.Offset(offset).Limit(perPage)
	if name != "" {
		if err := query.Where("user_id = ? AND name LIKE ?", userID, "%"+name+"%").Find(&campaigns).Error; err != nil {
			return nil, err
		}
		return campaigns, nil
	}

	if err := query.Where("user_id = ?", userID).Find(&campaigns).Error; err != nil {
		return nil, err
	}

	return campaigns, nil
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
	if err := r.DB.Model(&campaign).Updates(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

func NewCampaignRepository(DB *gorm.DB) campaign.CampaignRepositoryInterface {
	return &CampaignRepositoryImpl{
		DB: DB,
	}
}
