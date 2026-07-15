package database

import (
	"embed"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"hadirin-back/config"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func Migrate(cfg *config.Config) {
	source, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		log.Fatalf("Gagal membaca file migration: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", source, cfg.MigrationURL())
	if err != nil {
		log.Fatalf("Gagal menyiapkan migration: %v", err)
	}

	// migrate.ErrNoChange artinya semua migration sudah pernah dijalankan — bukan error
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Gagal menjalankan migration: %v", err)
	}

	log.Println("Migration database sudah paling baru")
}
