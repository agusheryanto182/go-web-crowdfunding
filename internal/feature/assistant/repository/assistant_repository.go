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
	panic("unimplemented")
}

// CreateQuestion implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) CreateQuestion(chat *entity.AssistantModel) error {
	panic("unimplemented")
}

// GetChatByIdUser implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) GetChatByIdUser(ID uint64) ([]*entity.AssistantModel, error) {
	panic("unimplemented")
}

// GetLastDonateByUserID implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) GetLastDonateByUserID(userID uint64) ([]*entity.TransactionModels, error) {
	panic("unimplemented")
}

// GetTopRatedCampaigns implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) GetTopRatedCampaigns() ([]*entity.CampaignModels, error) {
	panic("unimplemented")
}

// GetTrendDonateCampaigns implements assistant.RepositoryAssistantInterface.
func (r *RepositoryAssistantImpl) GetTrendDonateCampaigns() ([]string, error) {
	panic("unimplemented")
}

func NewAssistantRepository(db *gorm.DB) assistant.RepositoryAssistantInterface {
	return &RepositoryAssistantImpl{
		db: db,
	}
}
