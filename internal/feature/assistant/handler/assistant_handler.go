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
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "forbidden : you do not have permission")
	}

	chatRequest := new(dto.CreateChatRequest)
	if err := c.BodyParser(chatRequest); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	data := &entity.AssistantModel{
		Text: chatRequest.Text,
	}

	result, err := h.service.CreateAnswer(uint64(currentUser.ID), data)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create answer : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", result)
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

	data := &entity.AssistantModel{
		Text: chatRequest.Text,
	}

	err := h.service.CreateQuestion(uint64(currentUser.ID), data)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create question : "+err.Error())
	}

	answer, err := h.service.CreateAnswer(uint64(currentUser.ID), data)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create answer : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", answer)
}

// GenerateArticle implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) GenerateArticle(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role != entity.RoleAdmin {
		return response.SendStatusForbidden(c, "forbidden : you do not have permission to access this")
	}

	request := new(dto.GenerateArticleAiRequest)
	if err := c.BodyParser(&request); err != nil {
		return response.SendStatusBadRequest(c, "invalid input")
	}

	chat, err := h.service.GenerateArticle(request.Text)
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", chat)
}

// GetCampaignByUserID implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) GetCampaignByUserID(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)

	campaigns, err := h.service.GenerateRecommendationCampaign(uint64(currentUser.ID))
	if err != nil {
		return response.SendStatusNotFound(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", campaigns)
}

// GetChatByIdUser implements assistant.HandlerAssistantInterface.
func (h *HandlerAssistantImpl) GetChatByIdUser(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)

	chats, err := h.service.GetChatByUserID(uint64(currentUser.ID))
	if err != nil {
		return response.SendStatusNotFound(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", chats)
}

func NewHandlerAssistant(service assistant.ServiceAssistantInterface) assistant.HandlerAssistantInterface {
	return &HandlerAssistantImpl{
		service: service,
	}
}
