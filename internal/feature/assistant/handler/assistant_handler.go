package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
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
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "forbidden : you do not have permission")
	}

	chatRequest := new(dto.CreateChatRequest)
	if err := c.BodyParser(chatRequest); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	newUser := &entity.AssistantModel{
		Text: chatRequest.Text,
	}

	err := h.service.CreateQuestion(uint64(currentUser.ID), newUser)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create question : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", chatRequest.Text)
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
