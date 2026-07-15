package user

import (
	"github.com/gofiber/fiber/v2"

	"hadirin-back/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data user")
	}

	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data user", users)
}
