package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerImpl struct {
	userService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) user.UserHandlerInterface {
	return &UserHandlerImpl{userService: userService}
}

func (h *UserHandlerImpl) Update(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)

	updateUserRequest := dto.UpdateUserRequest{}
	if err := c.BodyParser(&updateUserRequest); err != nil {
		return response.SendStatusBadRequest(c, "invalid input")
	}

	if err := validator.ValidateStruct(&updateUserRequest); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	_, err := h.userService.UpdateUser(currentUser.ID, &updateUserRequest)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to update user : "+err.Error())
	}

	return response.SendStatusOkResponse(c, "success to update user")
}
