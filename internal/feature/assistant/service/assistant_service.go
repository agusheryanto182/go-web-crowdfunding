package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/config"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type ServiceAssistantImpl struct {
	repo   assistant.RepositoryAssistantInterface
	debug  bool
	openai *openai.Client
	config config.Config
}

// CreateAnswer implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) CreateAnswer(userID uint64, newData *entity.AssistantModel) (string, error) {
	ctx := context.Background()
	chat := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Menggunakan bahasa Indonesia. Kamu adalah Chatbot bertema crowdfunding",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Hi, bisa bantu saya menjawab tentang crowdfunding?",
		},
	}

	if newData.Text != "" {
		chat = append(chat, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Kamu akan diberikan sebuah pertanyaan mengenai %s, berikan jawabannya maksimal 20 kata", newData.Text),
		})
	}

	resp, err := s.GetAnswerFromAi(chat, ctx)
	if err != nil {
		logrus.Error("Can't Get Answer From Ai : " + err.Error())
	}

	if s.debug {
		fmt.Printf(
			"ID: %s. Created: %d. Model: %s. Choices: %v.\n",
			resp.ID, resp.Created, resp.Model, resp.Choices,
		)
	}

	answer := openai.ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	answerText := fmt.Sprintf(resp.Choices[0].Message.Content)

	value := &entity.AssistantModel{
		UserID:    userID,
		Role:      "answer",
		Text:      answer.Content,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateAnswer(value); err != nil {
		logrus.Error("Can't create answer in the repository: ", err.Error())
		return "", err
	}
	return answerText, nil
}

// CreateQuestion implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) CreateQuestion(userID uint64, newData *entity.AssistantModel) error {
	value := &entity.AssistantModel{
		UserID:    userID,
		Role:      "question",
		Text:      newData.Text,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateQuestion(value); err != nil {
		return err
	}
	return nil
}

// GenerateArticle implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GenerateArticle(title string) (string, error) {
	ctx := context.Background()
	chat := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Menggunakan bahasa Indonesia. Kamu adalah Chatbot bertema crowdfunding",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Hi, bisa bantu saya menjawab tentang crowdfunding?",
		},
	}

	if title != "" {
		chat = append(chat, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Buatlah Artikel dengan judul %s", title),
		})
	}

	resp, err := s.GetAnswerFromAi(chat, ctx)
	if err != nil {
		logrus.Error("Can't Get Answer From Ai : ", err.Error())
	}

	if s.debug {
		fmt.Printf(
			"ID: %s. Created: %d. Model: %s. Choices: %v.\n",
			resp.ID, resp.Created, resp.Model, resp.Choices,
		)
	}

	answer := openai.ChatCompletionMessage{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	return answer.Content, nil
}

// GenerateRecommendationCampaign implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GenerateRecommendationCampaign(userID uint64) ([]string, error) {
	ctx := context.Background()

	donates, err := s.repo.GetLastDonateByUserID(userID)
	if err != nil {
		return nil, err
	}

	var recommendedCampaigns []string

	if len(donates) == 0 {
		topCampaign, err := s.repo.GetTopDonatedCampaigns()
		if err != nil {
			return nil, err
		}
		recommendedCampaigns = append(recommendedCampaigns, topCampaign...)

	} else {
		chat := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "You are an analyst for the user's purchases.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Based on the user's previous purchases, provide 3 relevant products. Just the product names, no need for description or others.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "example answers\n1. Pendeteksi asap\n2. Robot pembersih udara\n3. Sistem analisis opini",
			},
		}

		donateContent := "List of products donated by the user : \n"
		for _, donate := range donates {
			donateContent += fmt.Sprintf("- %s\n", donate.Campaigns.Name)
		}

		chat = append(chat, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: donateContent,
		})

		resp, err := s.GetAnswerFromAi(chat, ctx)
		if err != nil {
			return nil, err
		}
		for _, choice := range resp.Choices {
			if choice.Message.Role == "assistant" {
				lines := strings.Split(choice.Message.Content, "\n")
				for _, line := range lines {
					if strings.TrimSpace(line) != "" {
						campaign := strings.SplitN(line, ". ", 2)
						if len(campaign) > 1 {
							recommendedCampaigns = append(recommendedCampaigns, strings.TrimSpace(campaign[1]))
						}
					}
				}
			}
		}
	}
	return recommendedCampaigns, nil
}

// GetAnswerFromAi implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GetAnswerFromAi(chat []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error) {
	model := openai.GPT3Dot5Turbo
	resp, err := s.openai.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: chat,
		},
	)
	return resp, err
}

// GetChatByUserID implements assistant.ServiceAssistantInterface.
func (s *ServiceAssistantImpl) GetChatByUserID(userID uint64) ([]*entity.AssistantModel, error) {
	panic("unimplemented")
}

func NewServiceAssistant(repo assistant.RepositoryAssistantInterface, openai *openai.Client, config config.Config) assistant.ServiceAssistantInterface {
	return &ServiceAssistantImpl{
		repo:   repo,
		debug:  false,
		openai: openai,
		config: config,
	}
}
