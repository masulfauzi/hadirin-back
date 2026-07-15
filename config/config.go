package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	CORSOrigins    string
	JWTSecret      string
	JWTExpireHours int
}

func Load() *Config {
	// Tidak fatal jika .env tidak ada (misalnya di server production
	// yang env-nya di-set langsung oleh sistem)
	_ = godotenv.Load()

	cfg := &Config{
		AppPort:        getEnv("APP_PORT", "3000"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", "hadirin"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		CORSOrigins:    getEnv("CORS_ORIGINS", "http://localhost:5173"),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpireHours: getEnvInt("JWT_EXPIRE_HOURS", 24),
	}

	// JWT_SECRET tidak boleh punya nilai default — aplikasi harus
	// menolak berjalan jika lupa di-set
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET wajib diisi di .env")
	}

	return cfg
}

// DSN menghasilkan connection string untuk PostgreSQL (dipakai GORM)
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

// MigrationURL menghasilkan connection string berformat URL (dipakai golang-migrate)
func (c *Config) MigrationURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, url.QueryEscape(c.DBPassword), c.DBHost, c.DBPort, c.DBName, c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if n, err := strconv.Atoi(value); err == nil {
			return n
		}
	}
	return fallback
}
