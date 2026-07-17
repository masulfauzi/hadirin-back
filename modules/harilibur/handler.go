package harilibur

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
	DivisionID *uuid.UUID `json:"division_id"`
	Tanggal    string     `json:"tanggal"`
	Keterangan string     `json:"keterangan"`
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	list, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data hari libur")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data hari libur", list)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	item, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Hari libur tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data hari libur", item)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	item := &HariLibur{DivisionID: req.DivisionID, Tanggal: req.Tanggal, Keterangan: req.Keterangan}
	if err := h.service.Create(item); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil membuat hari libur", item)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	item, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Hari libur tidak ditemukan")
	}

	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	item.DivisionID = req.DivisionID
	item.Tanggal = req.Tanggal
	item.Keterangan = req.Keterangan

	if err := h.service.Update(item); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengubah hari libur", item)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Hari libur tidak ditemukan")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus hari libur")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus hari libur", nil)
}
