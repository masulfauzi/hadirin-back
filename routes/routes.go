package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/config"
	"hadirin-back/modules/auth"
	"hadirin-back/modules/health"
	"hadirin-back/modules/user"
)

func Setup(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	api := app.Group("/api/v1")

	health.RegisterRoutes(api, db)

	// userService dibagi ke modul user dan auth
	userService := user.NewService(user.NewRepository(db))
	user.RegisterRoutes(api, userService, cfg)
	auth.RegisterRoutes(api, userService, cfg)

	// Modul berikutnya didaftarkan di sini, contoh:
	// presensi.RegisterRoutes(api, db, cfg)
	// rekap.RegisterRoutes(api, db, cfg)
	// ijin.RegisterRoutes(api, db, cfg)
}
