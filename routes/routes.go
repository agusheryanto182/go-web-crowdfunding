package routes

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/transaction"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/middleware"
	"github.com/agusheryanto182/go-web-crowdfunding/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, handler user.UserHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	userGroup := app.Group("api/v1/user")
	userGroup.Put("/edit-profile", middleware.Protected(jwtService, userService), handler.Update)
	userGroup.Post("/upload-avatar", middleware.Protected(jwtService, userService), handler.UploadAvatar)
	userGroup.Get("/", middleware.Protected(jwtService, userService), handler.GetAllUser)
	userGroup.Delete("/:id", middleware.Protected(jwtService, userService), handler.Delete)
}

func AuthRoute(app *fiber.App, handler auth.AuthHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	authGroup := app.Group("api/v1/auth")
	authGroup.Post("/sign-up", handler.SignUp)
	authGroup.Post("/verify-otp", handler.VerifyOTP)
	authGroup.Post("/sign-in", handler.SignIn)
}

func CampaignRoute(app *fiber.App, handler campaign.CampaignHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	campaignGroup := app.Group("api/v1/campaign")
	campaignGroup.Post("/save", middleware.Protected(jwtService, userService), handler.Save)
	campaignGroup.Get("/:id", handler.GetByID)
	campaignGroup.Get("/", handler.GetAll)
	campaignGroup.Put("/update/:id", middleware.Protected(jwtService, userService), handler.Update)
	campaignGroup.Post("/image/:id/upload", middleware.Protected(jwtService, userService), handler.CreateImage)
	campaignGroup.Put("/:id/image", middleware.Protected(jwtService, userService), handler.SetPrimaryImage)
	campaignGroup.Delete("/:id/image", middleware.Protected(jwtService, userService), handler.DeleteImageCampaign)
	campaignGroup.Delete("/:id", middleware.Protected(jwtService, userService), handler.DeleteCampaign)
}

func TransactionRoute(app *fiber.App, handler transaction.TransactionHandlerInterface, jwtService jwt.IJwt, userService user.UserServiceInterface) {
	transactionGroup := app.Group("api/v1/transaction")
	transactionGroup.Get("/", middleware.Protected(jwtService, userService), handler.GetAllCampaignTransactions)
	transactionGroup.Post("/", middleware.Protected(jwtService, userService), handler.CreateTransaction)
	transactionGroup.Post("/notifications", handler.GetNotification)
}
