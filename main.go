package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"hadirin-back/config"
	"hadirin-back/database"
	"hadirin-back/routes"
)

func main() {
	cfg := config.Load()

	// Jalankan migration dulu, baru buka koneksi GORM
	database.Migrate(cfg)
	db := database.Connect(cfg)

	app := fiber.New(fiber.Config{
		AppName: "Hadirin Backend",
	})

	// Logger: mencatat setiap request ke terminal
	app.Use(logger.New())

	// CORS: wajib agar frontend Vue.js (browser) bisa mengakses API
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigins,
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	routes.Setup(app, db, cfg)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
