package routes

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/auth"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/campaign"
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
}
