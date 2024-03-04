package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/dto"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/caching"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/email"
	utils "github.com/agusheryanto182/go-web-crowdfunding/utils/hash"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/otp"
)

type AuthServiceImpl struct {
	authRepo    auth.AuthRepositoryInterface
	userService user.UserServiceInterface
	hash        utils.HashInterface
	email       email.EmailSenderInterface
	cache       caching.CacheRepository
}

func NewAuthService(
	authRepo auth.AuthRepositoryInterface,
	userService user.UserServiceInterface,
	hash utils.HashInterface,
	email email.EmailSenderInterface,
	cache caching.CacheRepository,
) auth.AuthServiceInterface {
	return &AuthServiceImpl{
		authRepo:    authRepo,
		userService: userService,
		hash:        hash,
		email:       email,
		cache:       cache,
	}
}

func (s *AuthServiceImpl) SignUp(payload *dto.RegisterUserRequest) (*entity.UserModels, error) {
	isExistEmail, _ := s.userService.IsAvailableEmail(payload.Email)
	if isExistEmail {
		return nil, errors.New("email is already exist")
	}

	user := &entity.UserModels{
		Name:       payload.Name,
		Occupation: payload.Occupation,
		Email:      payload.Email,
		Role:       "user",
	}

	passwordHash, err := s.hash.GenerateHash(payload.Password)
	if err != nil {
		return &entity.UserModels{}, errors.New("failed to generate hash")
	}

	user.Password = passwordHash

	result, err := s.authRepo.SignUp(user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	generateOTP := otp.GenerateRandomOTP(7)
	newOTP := &entity.OTPModels{
		UserID:     int(result.ID),
		OTP:        generateOTP,
		ExpiredOTP: time.Now().Add(3 * time.Minute).Unix(),
	}

	if _, err := s.authRepo.SaveOTP(newOTP); err != nil {
		return nil, err
	}

	fmt.Println(result.Email)

	err = s.email.QueueEmail(result.Email, generateOTP)
	if err != nil {
		return nil, err
	}

	return result, nil
}
