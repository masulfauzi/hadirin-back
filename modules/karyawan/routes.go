package karyawan

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/config"
	"hadirin-back/middleware"
	"hadirin-back/modules/division"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config, divisionService *division.Service) *Service {
	service := NewService(NewRepository(db), divisionService)
	handler := NewHandler(service)

	g := router.Group("/karyawan", middleware.Protected(cfg.JWTSecret))
	g.Get("/", handler.GetAll)
	g.Get("/:id", handler.GetByID)
	g.Post("/", handler.Create)
	g.Put("/:id", handler.Update)
	g.Delete("/:id", handler.Delete)

	return service
}
