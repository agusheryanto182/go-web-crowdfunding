package main

import (
	"fmt"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/config"
	authHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/handler"
	authRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/repository"
	authService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/service"
	paymentService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/payment/service"
	"github.com/sashabaranov/go-openai"

	userHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/handler"
	userRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/repository"
	userService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/service"

	campaignHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/handler"
	campaignRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/repository"
	campaignService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/service"

	transactionHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/handler"
	transactionRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/repository"
	transactionService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/service"

	assistantHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant/handler"
	assistantRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant/repository"
	assistantService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/assistant/service"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/middleware"
	"github.com/agusheryanto182/go-web-crowdfunding/routes"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/caching/redis"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/database"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/email"
	utils "github.com/agusheryanto182/go-web-crowdfunding/utils/hash"
	Njwt "github.com/agusheryanto182/go-web-crowdfunding/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "API crowdfunding",
	})

	bootConfig := config.BootConfig()

	hash := utils.NewHash()
	jwt := Njwt.NewJWT(bootConfig.Secret)
	rdb := redis.InitialRedis(*bootConfig)
	cache := redis.NewRedisClient(rdb)
	mail := email.NewEmailService(rdb)

	DB := database.InitialDB(*bootConfig)

	paymentService := paymentService.NewPaymentService(*bootConfig)

	database.TableMigration(DB)

	// user
	userRepo := userRepo.NewUserRepository(DB)
	userService := userService.NewUserService(userRepo, hash)
	userHandler := userHandler.NewUserHandler(userService)

	// auth
	authRepo := authRepo.NewAuthRepository(DB)
	authService := authService.NewAuthService(authRepo, userRepo, userService, hash, mail, cache, jwt)
	authHandler := authHandler.NewAuthHandler(authService)

	// campaign
	campaignRepo := campaignRepo.NewCampaignRepository(DB)
	campaignService := campaignService.NewCampaignService(campaignRepo, userService)
	campaignHandler := campaignHandler.NewCampaignHandler(campaignService)

	// transaction
	transactionRepo := transactionRepo.NewTransactionRepository(DB)
	transactionService := transactionService.NewTransactionService(transactionRepo, paymentService, campaignRepo, campaignService)
	transactionHandler := transactionHandler.NewTransactionHandler(transactionService, campaignService)

	var client = openai.NewClient(bootConfig.OpenAiApiKey)
	// assistant
	assistantRepo := assistantRepo.NewAssistantRepository(DB)
	assistantService := assistantService.NewServiceAssistant(assistantRepo, client, *bootConfig)
	assistantHandler := assistantHandler.NewHandlerAssistant(assistantService)

	app.Use(middleware.Logging())

	// route
	routes.UserRoute(app, userHandler, jwt, userService)
	routes.AuthRoute(app, authHandler, jwt, userService)
	routes.CampaignRoute(app, campaignHandler, jwt, userService)
	routes.TransactionRoute(app, transactionHandler, jwt, userService)
	routes.AssistantRoutes(app, assistantHandler, jwt, userService)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)

	go mail.Worker()

	if err := app.Listen(addr).Error(); err != addr {
		panic("Application failed to start")
	}

}
