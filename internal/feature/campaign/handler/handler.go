package handler

import (
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/upload"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type CampaignHandlerImpl struct {
	service campaign.CampaignServiceInterface
}

// DeleteCampaign implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) DeleteCampaign(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)

	IdCampaign, _ := strconv.Atoi(c.Params("id"))

	campaign, err := h.service.GetByID(IdCampaign)
	if err != nil {
		return response.SendStatusNotFound(c, "not found : "+err.Error())
	}

	if campaign.UserID != currentUser.ID {
		return response.SendStatusForbidden(c, "forbidden : "+err.Error())
	}

	if err := h.service.DeleteCampaign(IdCampaign); err != nil {
		return response.SendStatusBadRequest(c, "failed : "+err.Error())
	}
	return response.SendStatusOkResponse(c, "success")
}

// DeleteImageCampaign implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) DeleteImageCampaign(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)

	IdCampaign, _ := strconv.Atoi(c.Params("id"))

	IdImage, _ := strconv.Atoi(c.Query("id"))

	if _, err := h.service.FindImageByID(IdImage); err != nil {
		return response.SendStatusNotFound(c, "error : "+err.Error())
	}

	campaign, err := h.service.GetByID(IdCampaign)
	if err != nil {
		return response.SendStatusNotFound(c, "error : "+err.Error())
	}

	if campaign.UserID != currentUser.ID {
		return response.SendStatusForbidden(c, "forbidden : you cannot access this request")
	}

	if err := h.service.DeleteImageCampaign(IdCampaign, IdImage); err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}
	return response.SendStatusOkResponse(c, "success")
}

// SetPrimaryImage implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) SetPrimaryImage(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)

	IdCampaign, _ := strconv.Atoi(c.Params("id"))

	IdImage, _ := strconv.Atoi(c.Query("id"))

	if _, err := h.service.FindImageByID(IdImage); err != nil {
		return response.SendStatusNotFound(c, "error : "+err.Error())
	}

	campaign, err := h.service.GetByID(IdCampaign)
	if err != nil {
		return response.SendStatusNotFound(c, "error : "+err.Error())
	}

	if campaign.UserID != currentUser.ID {
		return response.SendStatusForbidden(c, "forbidden : you cannot access this request")
	}

	payload := &dto.SetPrimaryImageRequest{
		ID:         IdImage,
		CampaignID: IdCampaign,
	}

	if _, err := h.service.SetPrimaryImage(payload); err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	return response.SendStatusOkResponse(c, "success")
}

// CreateImage implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) CreateImage(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	ID, _ := strconv.Atoi(c.Params("id"))

	setPrimaryImage := &dto.SetPrimaryImageRequest{}

	campaign, err := h.service.GetByID(ID)
	if campaign == nil {
		return response.SendStatusNotFound(c, "not found : "+err.Error())
	}
	if campaign.UserID != currentUser.ID {
		return response.SendStatusForbidden(c, "forbidden : "+err.Error())
	}

	payload := &dto.CreateRequestCampaignImage{}

	form, _ := c.MultipartForm()
	files := form.File["images"]
	for _, file := range files {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.SendStatusInternalServerError(c, "failed to open file : "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)

		imageURL, err := upload.ImageUploadHelper(fileToUpload)
		if err != nil {
			return response.SendStatusInternalServerError(c, "failed to upload image : "+err.Error())
		}

		payload.FileName = imageURL
		payload.CampaignID = ID

		result, err := h.service.CreateImage(payload)
		if err != nil {
			return response.SendStatusBadRequest(c, "error : "+err.Error())
		}
		setPrimaryImage.ID = result.ID
		setPrimaryImage.CampaignID = payload.CampaignID
	}

	if _, err := h.service.SetPrimaryImage(setPrimaryImage); err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}

	return response.SendStatusOkResponse(c, "success")
}

// GetAll implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page"))
	perPage := 5

	var campaign []*entity.CampaignModels
	var totalItems int64
	var err error
	search := c.Query("search")
	userID, _ := strconv.Atoi(c.Query("user_id"))

	fmt.Println(search)
	fmt.Println(userID)

	switch {
	case search != "" && userID == 0:
		campaign, totalItems, err = h.service.FindByNameWithPagination(page, perPage, search)
	case userID != 0 && search != "" || search == "":
		campaign, totalItems, err = h.service.GetByUserID(page, perPage, userID, search)
	default:
		campaign, totalItems, err = h.service.GetAll(page, perPage)
	}
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	currentPage, totalPages := h.service.CalculatePaginationValues(page, int(totalItems), perPage)
	nextPage := h.service.GetNextPage(currentPage, totalPages)
	prevPage := h.service.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.FormatCampaignsResponse(campaign), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success")

}

// GetByID implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) GetByID(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}
	campaign, err := h.service.GetByID(ID)
	if err != nil {
		return response.SendStatusNotFound(c, "error : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", *dto.FormatSaveCampaignResponse(campaign))
}

// Save implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) Save(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)
	if currentUser == nil {
		return response.SendStatusInternalServerError(c, "something error with middleware")
	}

	campaign := &dto.CreateRequestCampaign{}
	if err := c.BodyParser(campaign); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	if err := validator.ValidateStruct(campaign); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	campaign.UserID = currentUser.ID

	result, err := h.service.Save(campaign)
	if err != nil {
		return response.SendStatusBadRequest(c, "error : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", *dto.FormatSaveCampaignResponse(result))

}

// Update implements campaign.CampaignHandlerInterface.
func (h *CampaignHandlerImpl) Update(c *fiber.Ctx) error {
	currentUser := c.Locals("CurrentUser").(*entity.UserModels)

	ID, _ := strconv.Atoi(c.Params("id"))

	payload := &dto.UpdateRequestCampaign{}

	if err := c.BodyParser(payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid input : "+err.Error())
	}

	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	payload.ID = ID

	result, err := h.service.Update(currentUser.ID, payload)
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success", dto.FormatSaveCampaignResponse(result))
}

func NewCampaignHandler(service campaign.CampaignServiceInterface) campaign.CampaignHandlerInterface {
	return &CampaignHandlerImpl{
		service: service,
	}
}
