package menu

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
	ParentID *uuid.UUID `json:"parent_id"`
	KodeMenu string     `json:"kode_menu"`
	NamaMenu string     `json:"nama_menu"`
	Icon     *string    `json:"icon"`
	Route    *string    `json:"route"`
	Urutan   int        `json:"urutan"`
	IsActive *bool      `json:"is_active"`
}

type permissionRequest struct {
	RoleID    uuid.UUID `json:"role_id"`
	CanShow   bool      `json:"can_show"`
	CanRead   bool      `json:"can_read"`
	CanInsert bool      `json:"can_insert"`
	CanUpdate bool      `json:"can_update"`
	CanDelete bool      `json:"can_delete"`
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	menus, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data menu")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data menu", menus)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	m, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Menu tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data menu", m)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	m := &Menu{
		ParentID: req.ParentID,
		KodeMenu: req.KodeMenu,
		NamaMenu: req.NamaMenu,
		Icon:     req.Icon,
		Route:    req.Route,
		Urutan:   req.Urutan,
	}
	if req.IsActive != nil {
		m.IsActive = *req.IsActive
	} else {
		m.IsActive = true
	}

	if err := h.service.Create(m); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil membuat menu", m)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	m, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Menu tidak ditemukan")
	}

	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	m.ParentID = req.ParentID
	m.KodeMenu = req.KodeMenu
	m.NamaMenu = req.NamaMenu
	m.Icon = req.Icon
	m.Route = req.Route
	m.Urutan = req.Urutan
	if req.IsActive != nil {
		m.IsActive = *req.IsActive
	}

	if err := h.service.Update(m); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengubah menu", m)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Menu tidak ditemukan")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus menu")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus menu", nil)
}

func (h *Handler) GetPermissions(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	perms, err := h.service.GetPermissions(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Menu tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil permission menu", perms)
}

func (h *Handler) SetPermission(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	var req permissionRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	p := &RoleMenuPermission{
		RoleID:    req.RoleID,
		CanShow:   req.CanShow,
		CanRead:   req.CanRead,
		CanInsert: req.CanInsert,
		CanUpdate: req.CanUpdate,
		CanDelete: req.CanDelete,
	}

	if err := h.service.SetPermission(id, p); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menyimpan permission menu", p)
}
