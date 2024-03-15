package handler

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/response"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthHandlerImpl struct {
	authService auth.AuthServiceInterface
}

func NewAuthHandler(authService auth.AuthServiceInterface) auth.AuthHandlerInterface {
	return &AuthHandlerImpl{
		authService: authService,
	}
}

func (h *AuthHandlerImpl) SignUp(c *fiber.Ctx) error {

	var input dto.RegisterUserRequest
	if err := c.BodyParser(&input); err != nil {
		return response.SendStatusBadRequest(c, "invalid input")
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	result, err := h.authService.SignUp(&input)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to register user : "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success to create user", dto.FormatCreateUserResponse(result))
}

func (h *AuthHandlerImpl) VerifyOTP(c *fiber.Ctx) error {
	var requestVerify dto.VerifyOTPRequest
	if err := c.BodyParser(&requestVerify); err != nil {
		return response.SendStatusBadRequest(c, "invalid input")
	}

	if err := validator.ValidateStruct(&requestVerify); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	accessToken, err := h.authService.VerifyOTP(requestVerify.Email, requestVerify.OTP)
	if err != nil {
		return response.SendStatusBadRequest(c, "verify otp is failed : "+err.Error())
	}

	result := dto.VerifyOTPResponse{
		AccessToken: accessToken,
	}

	return response.SendStatusOkWithDataResponse(c, "verification is success", result)
}

func (h *AuthHandlerImpl) SignIn(c *fiber.Ctx) error {
	requestSignIn := &dto.SignInUserRequest{}

	if err := c.BodyParser(requestSignIn); err != nil {
		return response.SendStatusBadRequest(c, "invalid input")
	}

	if err := validator.ValidateStruct(requestSignIn); err != nil {
		return response.SendStatusBadRequest(c, "validation error : "+err.Error())
	}

	accessToken, err := h.authService.SignIn(requestSignIn)
	if err != nil {
		return response.SendStatusForbidden(c, "failed SignIn : "+err.Error())
	}

	return response.SendStatusCreatedWithDataResponse(c, "success", accessToken)
}
