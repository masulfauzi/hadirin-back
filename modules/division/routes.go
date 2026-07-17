package division

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/config"
	"hadirin-back/middleware"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config) *Service {
	service := NewService(NewRepository(db))
	handler := NewHandler(service)

	g := router.Group("/divisions", middleware.Protected(cfg.JWTSecret))
	g.Get("/", handler.GetAll)
	g.Get("/:id", handler.GetByID)
	g.Post("/", handler.Create)
	g.Put("/:id", handler.Update)
	g.Delete("/:id", handler.Delete)
	g.Get("/:id/schedules", handler.GetSchedules)
	g.Put("/:id/schedules", handler.SetSchedules)

	return service
}
