package auth

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

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}
	if req.Name == "" || req.Email == "" {
		return utils.Error(c, fiber.StatusBadRequest, "name dan email wajib diisi")
	}
	if len(req.Password) < 8 {
		return utils.Error(c, fiber.StatusBadRequest, "password minimal 8 karakter")
	}

	newUser, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return utils.Error(c, fiber.StatusConflict, "Email sudah terdaftar")
	}

	return utils.Success(c, fiber.StatusCreated, "Registrasi berhasil", newUser)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	token, loggedUser, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Login berhasil", fiber.Map{
		"token": token,
		"user":  loggedUser,
	})
}

func (h *Handler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	currentUser, err := h.service.GetProfile(userID)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "User tidak ditemukan")
	}

	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil profil", currentUser)
}
