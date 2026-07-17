package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"hadirin-back/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type upsertRequest struct {
	KodeIdentitas string  `json:"kode_identitas"`
	Username      string  `json:"username"`
	Email         *string `json:"email"`
	Password      string  `json:"password"`
	IsActive      *bool   `json:"is_active"`
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data user")
	}

	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data user", users)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	u, err := h.service.GetUserByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "User tidak ditemukan")
	}

	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data user", u)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetUserByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "User tidak ditemukan")
	}

	if err := h.service.DeleteUser(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus user")
	}

	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus user", nil)
}
