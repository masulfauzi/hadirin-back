package health

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	handler := NewHandler(db)

	router.Get("/health", handler.Check)
}
