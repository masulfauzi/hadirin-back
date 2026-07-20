package auth

import (
	"github.com/gofiber/fiber/v2"

	"hadirin-back/config"
	"hadirin-back/middleware"
	"hadirin-back/modules/karyawan"
	"hadirin-back/modules/role"
	"hadirin-back/modules/user"
)

func RegisterRoutes(router fiber.Router, userService *user.Service, karyawanService *karyawan.Service, roleService *role.Service, cfg *config.Config) {
	service := NewService(userService, karyawanService, roleService, cfg)
	handler := NewHandler(service)

	authGroup := router.Group("/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
	authGroup.Get("/me", middleware.Protected(cfg.JWTSecret), handler.Me)
}
