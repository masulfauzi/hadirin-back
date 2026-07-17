package role

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
	KodeRole  string  `json:"kode_role"`
	NamaRole  string  `json:"nama_role"`
	Deskripsi *string `json:"deskripsi"`
	IsActive  *bool   `json:"is_active"`
}

type assignUserRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	roles, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data role")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data role", roles)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	r, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Role tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data role", r)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	r := &Role{KodeRole: req.KodeRole, NamaRole: req.NamaRole, Deskripsi: req.Deskripsi}
	if req.IsActive != nil {
		r.IsActive = *req.IsActive
	} else {
		r.IsActive = true
	}

	if err := h.service.Create(r); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil membuat role", r)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	r, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Role tidak ditemukan")
	}

	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	r.KodeRole = req.KodeRole
	r.NamaRole = req.NamaRole
	r.Deskripsi = req.Deskripsi
	if req.IsActive != nil {
		r.IsActive = *req.IsActive
	}

	if err := h.service.Update(r); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengubah role", r)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Role tidak ditemukan")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus role")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus role", nil)
}

func (h *Handler) AssignUser(c *fiber.Ctx) error {
	roleID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	var req assignUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	if err := h.service.AssignUser(roleID, req.UserID); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil menetapkan role ke user", nil)
}

func (h *Handler) RevokeUser(c *fiber.Ctx) error {
	roleID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}
	userID, err := uuid.Parse(c.Params("userId"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "user ID tidak valid")
	}

	if err := h.service.RevokeUser(roleID, userID); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mencabut role dari user")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mencabut role dari user", nil)
}
