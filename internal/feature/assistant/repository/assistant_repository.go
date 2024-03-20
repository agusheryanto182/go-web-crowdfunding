package repository

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant"
	"gorm.io/gorm"
)

type RepositoryAssistantImpl struct {
	db *gorm.DB
}

// CreateAnswer implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) CreateAnswer(chat *entity.AssistantModel) error {
	if err := r.db.Create(&chat).Error; err != nil {
		return err
	}
	return nil
}

// CreateQuestion implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) CreateQuestion(chat *entity.AssistantModel) error {
	if err := r.db.Create(&chat).Error; err != nil {
		return err
	}
	return nil
}

// GetChatByIdUser implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) GetChatByIdUser(ID uint64) ([]*entity.AssistantModel, error) {
	var chats []*entity.AssistantModel

	if err := r.db.Where("user_id = ?", ID).Find(&chats).Error; err != nil {
		return nil, err
	}

	return chats, nil
}

// GetLastDonateByUserID implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) GetLastDonateByUserID(userID uint64) ([]*entity.TransactionModels, error) {
	var transactions []*entity.TransactionModels

	if err := r.db.Debug().
		Preload("Campaigns").
		Where("user_id = ? AND status = ?", userID, "paid").
		Limit(10).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTopRatedCampaigns implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) GetTopDonatedCampaigns() ([]string, error) {
	var topCampaigns []string

	if err := r.db.Table("campaigns").
		Select("name").
		Order("backer_count DESC").
		Limit(3).
		Pluck("name", &topCampaigns).
		Error; err != nil {
		return nil, err
	}
	return topCampaigns, nil
}

func NewAssistantRepository(db *gorm.DB) assistant.RepositoryAssistantInterface {
	return &RepositoryAssistantImpl{
		db: db,
	}
}
