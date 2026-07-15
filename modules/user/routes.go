package user

import (
	"github.com/gofiber/fiber/v2"

	"hadirin-back/config"
	"hadirin-back/middleware"
)

func RegisterRoutes(router fiber.Router, service *Service, cfg *config.Config) {
	handler := NewHandler(service)

	// Group ini hanya bisa diakses dengan token JWT yang valid
	users := router.Group("/users", middleware.Protected(cfg.JWTSecret))
	users.Get("/", handler.GetUsers)
}
