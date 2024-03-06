package service

import (
	"testing"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/mocks"
	userRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/mocks"
	user "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/service"
	"github.com/stretchr/testify/assert"

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

	t.Run("Success case", func(t *testing.T) {
		actual := generateCacheKey(email, action)
		assert.Equal(t, expected, actual)
	})
}
