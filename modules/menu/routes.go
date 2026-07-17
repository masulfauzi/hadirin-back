package menu

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/config"
	"hadirin-back/middleware"
	"hadirin-back/modules/role"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config, roleService *role.Service) {
	service := NewService(NewRepository(db), roleService)
	handler := NewHandler(service)

	g := router.Group("/menus", middleware.Protected(cfg.JWTSecret))
	g.Get("/", handler.GetAll)
	g.Get("/:id", handler.GetByID)
	g.Post("/", handler.Create)
	g.Put("/:id", handler.Update)
	g.Delete("/:id", handler.Delete)
	g.Get("/:id/permissions", handler.GetPermissions)
	g.Put("/:id/permissions", handler.SetPermission)
}
