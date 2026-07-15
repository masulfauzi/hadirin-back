package auth

import (
	"github.com/gofiber/fiber/v2"

	"hadirin-back/config"
	"hadirin-back/middleware"
	"hadirin-back/modules/user"
)

func RegisterRoutes(router fiber.Router, userService *user.Service, cfg *config.Config) {
	service := NewService(userService, cfg)
	handler := NewHandler(service)

	authGroup := router.Group("/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
	authGroup.Get("/me", middleware.Protected(cfg.JWTSecret), handler.Me)
}
