package main

import (
	"fmt"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/config"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/feature/user/repository"
	"github.com/agusheryanto182/go-web-crowdfunding/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "API crowdfunding",
	})

	bootConfig := config.NewConfig()

	DB := config.InitialDB(*bootConfig)

	config.TableMigration(DB)

	userRepo := repository.NewUserRepository(DB)

	app.Use(middleware.Logging())

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	if err := app.Listen(addr).Error(); err != addr {
		panic("Application failed to start")
	}
}
