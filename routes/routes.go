package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/config"
	"hadirin-back/modules/auth"
	"hadirin-back/modules/division"
	"hadirin-back/modules/harilibur"
	"hadirin-back/modules/health"
	"hadirin-back/modules/ijin"
	"hadirin-back/modules/karyawan"
	"hadirin-back/modules/menu"
	"hadirin-back/modules/presensi"
	"hadirin-back/modules/role"
	"hadirin-back/modules/user"
)

func Setup(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	api := app.Group("/api/v1")

	health.RegisterRoutes(api, db)

	// userService dibagi ke modul user dan auth
	userService := user.NewService(user.NewRepository(db))
	user.RegisterRoutes(api, userService, cfg)

	// division, karyawan & role didaftarkan lebih dulu karena service-nya
	// dipakai (di-inject) oleh modul lain, termasuk auth (register user +
	// karyawan + assign role karyawan sekaligus)
	divisionService := division.RegisterRoutes(api, db, cfg)
	karyawanService := karyawan.RegisterRoutes(api, db, cfg, divisionService)
	roleService := role.RegisterRoutes(api, db, cfg, userService)

	auth.RegisterRoutes(api, userService, karyawanService, roleService, cfg)

	menu.RegisterRoutes(api, db, cfg, roleService)
	harilibur.RegisterRoutes(api, db, cfg, divisionService)
	presensi.RegisterRoutes(api, db, cfg, karyawanService)
	ijin.RegisterRoutes(api, db, cfg, karyawanService)
}
