package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"hadirin-back/modules/role"
	"hadirin-back/modules/user"
	"hadirin-back/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type registerRequest struct {
	KodeIdentitas string  `json:"kode_identitas"`
	Username      string  `json:"username"`
	Email         *string `json:"email"`
	Password      string  `json:"password"`
	NamaLengkap   string  `json:"nama_lengkap"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type meResponse struct {
	user.User
	Roles []role.Role `json:"roles"`
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}
	if req.KodeIdentitas == "" || req.Username == "" || req.NamaLengkap == "" {
		return utils.Error(c, fiber.StatusBadRequest, "kode_identitas, username, dan nama_lengkap wajib diisi")
	}
	if len(req.Password) < 8 {
		return utils.Error(c, fiber.StatusBadRequest, "password minimal 8 karakter")
	}

	newUser, err := h.service.Register(req.KodeIdentitas, req.Username, req.Email, req.Password, req.NamaLengkap)
	if err != nil {
		return utils.Error(c, fiber.StatusConflict, err.Error())
	}

	return utils.Success(c, fiber.StatusCreated, "Registrasi berhasil", newUser)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	token, loggedUser, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Login berhasil", fiber.Map{
		"token": token,
		"user":  loggedUser,
	})
}

func (h *Handler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)

	currentUser, roles, err := h.service.GetProfile(userID)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "User tidak ditemukan")
	}

	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil profil", meResponse{User: *currentUser, Roles: roles})
}
