package assistant

import (
	"context"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
)

type RepositoryAssistantInterface interface {
	GetChatByIdUser(ID uint64) ([]*entity.AssistantModel, error)
	CreateQuestion(chat *entity.AssistantModel) error
	CreateAnswer(chat *entity.AssistantModel) error
	GetLastDonateByUserID(userID uint64) ([]*entity.TransactionModels, error)
	GetTopRatedCampaigns() ([]*entity.CampaignModels, error)
	GetTrendDonateCampaigns() ([]string, error)
}

type ServiceAssistantInterface interface {
	GetChatByUserID(userID uint64) ([]*entity.AssistantModel, error)
	CreateQuestion(userID uint64, newData *entity.AssistantModel) error
	CreateAnswer(userID uint64, newData *entity.AssistantModel) (string, error)
	GetAnswerFromAi(chat []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error)
	GenerateArticle(title string) (string, error)
	GenerateRecommendationCampaign(userID uint64) ([]string, error)
}

type HandlerAssistantInterface interface {
	GetChatByIdUser(c *fiber.Ctx) error
	CreateQuestion(c *fiber.Ctx) error
	CreateAnswer(c *fiber.Ctx) error
	GenerateArticle(c *fiber.Ctx) error
	GetCampaignByUserID(c *fiber.Ctx) error
}
