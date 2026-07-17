package presensi

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/config"
	"hadirin-back/middleware"
	"hadirin-back/modules/karyawan"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config, karyawanService *karyawan.Service) {
	service := NewService(NewRepository(db), karyawanService)
	handler := NewHandler(service)

	g := router.Group("/presensi", middleware.Protected(cfg.JWTSecret))
	g.Get("/", handler.GetAll)
	g.Get("/:id", handler.GetByID)
	g.Post("/", handler.Create)
	g.Put("/:id", handler.Update)
	g.Delete("/:id", handler.Delete)
}
