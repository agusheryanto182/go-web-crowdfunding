package main

import (
	"fmt"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/config"
	authHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/handler"
	authRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/repository"
	authService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth/service"
	paymentService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/payment/service"

	userHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/handler"
	userRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/repository"
	userService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/service"

	campaignHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/handler"
	campaignRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/repository"
	campaignService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign/service"

	transactionHandler "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/handler"
	transactionRepo "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/repository"
	transactionService "github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction/service"

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

	userRepo := userRepo.NewUserRepository(DB)
	userService := userService.NewUserService(userRepo, hash)
	userHandler := userHandler.NewUserHandler(userService)

	authRepo := authRepo.NewAuthRepository(DB)
	authService := authService.NewAuthService(authRepo, userRepo, userService, hash, mail, cache, jwt)
	authHandler := authHandler.NewAuthHandler(authService)

	campaignRepo := campaignRepo.NewCampaignRepository(DB)
	campaignService := campaignService.NewCampaignService(campaignRepo, userService)
	campaignHandler := campaignHandler.NewCampaignHandler(campaignService)

	transactionRepo := transactionRepo.NewTransactionRepository(DB)
	transactionService := transactionService.NewTransactionService(transactionRepo, paymentService, campaignRepo)
	transactionHandler := transactionHandler.NewTransactionHandler(transactionService, campaignService)

	app.Use(middleware.Logging())

	routes.UserRoute(app, userHandler, jwt, userService)
	routes.AuthRoute(app, authHandler, jwt, userService)
	routes.CampaignRoute(app, campaignHandler, jwt, userService)
	routes.TransactionRoute(app, transactionHandler, jwt, userService)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)

	go mail.Worker()

	if err := app.Listen(addr).Error(); err != addr {
		panic("Application failed to start")
	}

}
