package role

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/config"
	"hadirin-back/middleware"
	"hadirin-back/modules/user"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config, userService *user.Service) *Service {
	service := NewService(NewRepository(db), userService)
	handler := NewHandler(service)

	g := router.Group("/roles", middleware.Protected(cfg.JWTSecret))
	g.Get("/", handler.GetAll)
	g.Get("/:id", handler.GetByID)
	g.Post("/", handler.Create)
	g.Put("/:id", handler.Update)
	g.Delete("/:id", handler.Delete)
	g.Post("/:id/users", handler.AssignUser)
	g.Delete("/:id/users/:userId", handler.RevokeUser)

	return service
}
