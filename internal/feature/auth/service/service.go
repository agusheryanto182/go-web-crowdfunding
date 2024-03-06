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
	"github.com/agusheryanto182/go-web-crowdfunding/utils/jwt"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/otp"
)

type AuthServiceImpl struct {
	authRepo    auth.AuthRepositoryInterface
	userRepo    user.UserRepositoryInterface
	userService user.UserServiceInterface
	hash        utils.HashInterface
	email       email.EmailSenderInterface
	cache       caching.CacheRepository
	nJWT        jwt.IJwt
}

func NewAuthService(
	authRepo auth.AuthRepositoryInterface,
	userRepo user.UserRepositoryInterface,
	userService user.UserServiceInterface,
	hash utils.HashInterface,
	email email.EmailSenderInterface,
	cache caching.CacheRepository,
	nJWT jwt.IJwt,
) auth.AuthServiceInterface {
	return &AuthServiceImpl{
		authRepo:    authRepo,
		userRepo:    userRepo,
		userService: userService,
		hash:        hash,
		email:       email,
		cache:       cache,
		nJWT:        nJWT,
	}
}

func generateCacheKey(email, action string) string {
	return fmt.Sprintf("auth:%s:%s", email, action)
}

func (s *AuthServiceImpl) SignUp(payload *dto.RegisterUserRequest) (*entity.UserModels, error) {
	userByEmail, _ := s.userService.GetUserByEmail(payload.Email)
	if userByEmail != nil {
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
		return nil, errors.New("failed to generate hash")
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
		return nil, errors.New("failed to save OTP")
	}

	err = s.email.QueueEmail(result.Email, generateOTP)
	if err != nil {
		return nil, errors.New("failed add email to queue")
	}

	return result, nil
}

func (s *AuthServiceImpl) VerifyOTP(email, OTP string) (string, error) {
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if user.IsVerified {
		return "", errors.New("your account has been verified")
	}

	isValidOTP, err := s.authRepo.FindValidOTP(int(user.ID), OTP)
	if err != nil {
		return "", errors.New("invalid otp : " + err.Error())
	}

	user.IsVerified = true

	if _, err := s.userRepo.UpdateUser(user); err != nil {
		return "", errors.New("failed to update user : " + err.Error())
	}

	if err := s.authRepo.DeleteOTP(isValidOTP); err != nil {
		return "", errors.New("failed to delete otp : " + err.Error())
	}

	accessToken, err := s.nJWT.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return "", errors.New("failed to generate jwt : " + err.Error())
	}

	return accessToken, nil
}

func (s *AuthServiceImpl) SignIn(payload *dto.SignInUserRequest) (string, error) {
	check, err := s.userService.GetUserByEmail(payload.Email)
	if err != nil {
		return "", err
	}

	if !check.IsVerified {
		return "", errors.New("account has not been verified")
	}

	validPassword, err := s.hash.ComparePassword(check.Password, payload.Password)
	if err != nil {
		return "", err
	}

	if !validPassword {
		return "", errors.New("invalid credentials")
	}

	accessToken, err := s.nJWT.GenerateJWT(check.ID, check.Email, check.Role)
	if err != nil {
		return "", errors.New("failed to generate jwt")
	}

	return accessToken, nil
}
