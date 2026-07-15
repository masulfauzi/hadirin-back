package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hadirin-back/config"
)

func Connect(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	log.Println("Berhasil terhubung ke database")
	return db
}
