package service

import (
	"errors"
	"testing"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/mocks"
	userRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/mocks"
	user "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	utils "github.com/agusheryanto182/go-web-crowdfunding/utils/mocks"
)

func setupTestService(t *testing.T) (
	*AuthServiceImpl,
	*mocks.AuthRepositoryInterface,
	*utils.IJwt,
	*userRepo.UserRepositoryInterface,
	*utils.HashInterface,
	*utils.CacheRepository,
	*utils.EmailSenderInterface,
) {
	repo := mocks.NewAuthRepositoryInterface(t)
	jwt := utils.NewIJwt(t)
	userRepo := userRepo.NewUserRepositoryInterface(t)
	hash := utils.NewHashInterface(t)
	cache := utils.NewCacheRepository(t)
	email := utils.NewEmailSenderInterface(t)
	userService := user.NewUserService(userRepo, hash)
	service := NewAuthService(repo, userRepo, userService, hash, email, cache, jwt)

	return service.(*AuthServiceImpl), repo, jwt, userRepo, hash, cache, email
}

func Test_GenerateCache(t *testing.T) {
	email := "testing@gmail.com"
	action := "verify_status"
	expected := "auth:testing@gmail.com:verify_status"

	t.Run("success case", func(t *testing.T) {
		actual := generateCacheKey(email, action)
		assert.Equal(t, expected, actual)
		assert.NotNil(t, actual)
	})
}

func Test_SignUp(t *testing.T) {
	PasswordHash := "testingpassword"
	request := &dto.RegisterUserRequest{
		Email:    "testing@gmail.com",
		Password: "testing123",
	}

	existing := &entity.UserModels{
		Email:    "testing@gmail.com",
		Password: "testing123",
	}

	otp := &entity.OTPModels{
		OTP:        "1234567",
		UserID:     1,
		ExpiredOTP: 123234234,
	}

	expectedError := errors.New("some error")

	t.Run("failed case - email is already exist", func(t *testing.T) {
		authService, authRepo, _, userRepo, _, _, _ := setupTestService(t)
		userRepo.On("FindUserByEmail", request.Email).Return(existing, nil)
		result, err := authService.SignUp(request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "email is already exist")

		userRepo.AssertExpectations(t)
		authRepo.AssertExpectations(t)
	})

	t.Run("failed case - failed to generate hash", func(t *testing.T) {
		authService, authRepo, _, userRepo, hash, _, _ := setupTestService(t)
		userRepo.On("FindUserByEmail", request.Email).Return(nil, nil)
		hash.On("GenerateHash", request.Password).Return("", expectedError)
		result, err := authService.SignUp(request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "failed to generate hash")

		userRepo.AssertExpectations(t)
		authRepo.AssertExpectations(t)
	})

	t.Run("failed case - failed to create user", func(t *testing.T) {
		authService, authRepo, _, userRepo, hash, _, email := setupTestService(t)
		userRepo.On("FindUserByEmail", request.Email).Return(nil, nil)
		hash.On("GenerateHash", request.Password).Return(PasswordHash, nil)
		authRepo.On("SignUp", mock.AnythingOfType("*entity.UserModels")).Return(nil, expectedError)

		result, err := authService.SignUp(request)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "failed to create user")

		userRepo.AssertExpectations(t)
		authRepo.AssertExpectations(t)
		hash.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("failed case - failed to save OTP", func(t *testing.T) {
		authService, authRepo, _, userRepo, hash, _, email := setupTestService(t)
		userRepo.On("FindUserByEmail", request.Email).Return(nil, nil)
		hash.On("GenerateHash", request.Password).Return(PasswordHash, nil)
		authRepo.On("SignUp", mock.AnythingOfType("*entity.UserModels")).Return(existing, nil)
		authRepo.On("SaveOTP", mock.AnythingOfType("*entity.OTPModels")).Return(nil, expectedError)

		result, err := authService.SignUp(request)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "failed to save OTP")

		userRepo.AssertExpectations(t)
		authRepo.AssertExpectations(t)
		hash.AssertExpectations(t)
		email.AssertExpectations(t)
	})

	t.Run("failed case - failed to save OTP", func(t *testing.T) {
		authService, authRepo, _, userRepo, hash, _, email := setupTestService(t)
		userRepo.On("FindUserByEmail", request.Email).Return(nil, nil)
		hash.On("GenerateHash", request.Password).Return(PasswordHash, nil)
		authRepo.On("SignUp", mock.AnythingOfType("*entity.UserModels")).Return(existing, nil)
		authRepo.On("SaveOTP", mock.AnythingOfType("*entity.OTPModels")).Return(otp, nil)
		email.On("QueueEmail", otp.OTP).Return(expectedError)

		result, err := authService.SignUp(request)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "failed add email to queue")

		userRepo.AssertExpectations(t)
		authRepo.AssertExpectations(t)
		hash.AssertExpectations(t)
		email.AssertExpectations(t)
	})
}
