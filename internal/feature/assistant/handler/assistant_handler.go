package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant"
	"github.com/gofiber/fiber/v2"
)

type HandlerAssistantImpl struct {
	service assistant.ServiceAssistantInterface
}

// CreateAnswer implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) CreateAnswer(c *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateQuestion implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) CreateQuestion(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GenerateArticle implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) GenerateArticle(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetCampaignByUserID implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) GetCampaignByUserID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetChatByIdUser implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) GetChatByIdUser(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewHandlerAssistant(service assistant.ServiceAssistantInterface) assistant.HandlerAssistantInterface {
	return &HandlerAssistantImpl{
		service: service,
	}
}
