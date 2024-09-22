package controller

import (
	"myapi/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	actionPrefix  string
	authService *service.AuthService
}

func NewAuthController(app *fiber.App, authService *service.AuthService) {
	authController := &AuthController{
		actionPrefix:  `/auth`,
		authService: authService,
	}

	app.Route(authController.actionPrefix, func(router fiber.Router) {
	})
}