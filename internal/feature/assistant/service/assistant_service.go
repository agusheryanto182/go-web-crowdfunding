package service

import (
	"context"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant"
	openai "github.com/sashabaranov/go-openai"
)

type ServiceAssistantImpl struct {
	repo assistant.RepositoryAssistantInterface
}

// CreateAnswer implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) CreateAnswer(userID uint64, newData *entity.AssistantModel) (string, error) {
	panic("unimplemented")
}

// CreateQuestion implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) CreateQuestion(userID uint64, newData *entity.AssistantModel) error {
	panic("unimplemented")
}

// GenerateArticle implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GenerateArticle(title string) (string, error) {
	panic("unimplemented")
}

// GenerateRecommendationCampaign implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GenerateRecommendationCampaign(userID uint64) ([]string, error) {
	panic("unimplemented")
}

// GetAnswerFromAi implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GetAnswerFromAi(chat []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error) {
	panic("unimplemented")
}

// GetChatByUserID implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GetChatByUserID(userID uint64) ([]*entity.AssistantModel, error) {
	panic("unimplemented")
}

func NewServiceAssistant(repo assistant.RepositoryAssistantInterface) assistant.ServiceAssistantInterface {
	return &ServiceAssistantImpl{
		repo: repo,
	}
}
