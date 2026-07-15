package health

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"hadirin-back/utils"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Check(c *fiber.Ctx) error {
	dbStatus := "ok"
	sqlDB, err := h.db.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "error"
	}

	return utils.Success(c, fiber.StatusOK, "Server berjalan", fiber.Map{
		"status":   "ok",
		"database": dbStatus,
	})
}
