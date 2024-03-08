package handler

import (
	"mime/multipart"
	"strconv"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/upload"
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

func (h *UserHandlerImpl) UploadAvatar(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("CurrentUser").(*entity.UserModels)
	if !ok || currentUser == nil {
		return response.SendStatusUnauthorized(c, "user not found")
	}

	userUploadAvatar := &dto.UpdateAvatarRequest{}

	file, err := c.FormFile("avatar")
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.SendStatusInternalServerError(c, "failed to open file : "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)

		uploadedURL, err := upload.ImageUploadHelper(fileToUpload)
		if err != nil {
			return response.SendStatusInternalServerError(c, "failed to upload image : "+err.Error())
		}
		userUploadAvatar.Avatar = uploadedURL
	}

	image, err := h.userService.UploadAvatar(currentUser.ID, userUploadAvatar)
	if err != nil {
		return response.SendStatusBadRequest(c, "error upload avatar : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success upload avatar", dto.UpdateAvatarResponse(image))
}

func (h *UserHandlerImpl) GetAllUser(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser.Role != "admin" {
		return response.SendStatusForbidden(c, "forbidden : you dont have a permission for access this")
	}

	page, _ := strconv.Atoi(c.Query("page"))
	perPage := 8

	var user []*entity.UserModels
	var totalItems int64
	var err error
	search := c.Query("search")

	switch {
	case search != "":
		{
			user, totalItems, err = h.userService.GetUserByName(page, perPage, search)
		}
	default:
		user, totalItems, err = h.userService.GetAllUser(page, perPage)
	}
	if err != nil {
		c.Context().Logger()
		return response.SendStatusInternalServerError(c, "failed to get data user : "+err.Error())
	}

	currentPage, totalPages := h.userService.CalculatePaginationValues(page, int(totalItems), perPage)
	nextPage := h.userService.GetNextPage(currentPage, totalPages)
	prevPage := h.userService.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.FormatterUsers(user), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success")

}
