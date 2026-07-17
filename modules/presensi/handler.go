package presensi

import (
	"time"

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
	KaryawanID uuid.UUID `json:"karyawan_id"`
	Tanggal    string    `json:"tanggal"`

	JamHadir        *time.Time `json:"jam_hadir"`
	FotoHadir       *string    `json:"foto_hadir"`
	LatHadir        *float64   `json:"lat_hadir"`
	LongHadir       *float64   `json:"long_hadir"`
	JarakHadirMeter *float64   `json:"jarak_hadir_meter"`
	StatusHadir     *string    `json:"status_hadir"`

	JamPulang        *time.Time `json:"jam_pulang"`
	FotoPulang       *string    `json:"foto_pulang"`
	LatPulang        *float64   `json:"lat_pulang"`
	LongPulang       *float64   `json:"long_pulang"`
	JarakPulangMeter *float64   `json:"jarak_pulang_meter"`
	StatusPulang     *string    `json:"status_pulang"`
}

func toModel(req upsertRequest) *Presensi {
	return &Presensi{
		KaryawanID:       req.KaryawanID,
		Tanggal:          req.Tanggal,
		JamHadir:         req.JamHadir,
		FotoHadir:        req.FotoHadir,
		LatHadir:         req.LatHadir,
		LongHadir:        req.LongHadir,
		JarakHadirMeter:  req.JarakHadirMeter,
		StatusHadir:      req.StatusHadir,
		JamPulang:        req.JamPulang,
		FotoPulang:       req.FotoPulang,
		LatPulang:        req.LatPulang,
		LongPulang:       req.LongPulang,
		JarakPulangMeter: req.JarakPulangMeter,
		StatusPulang:     req.StatusPulang,
	}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	list, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data presensi")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data presensi", list)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	p, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Presensi tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data presensi", p)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	p := toModel(req)
	if err := h.service.Create(p); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil membuat presensi", p)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	existing, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Presensi tidak ditemukan")
	}

	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	updated := toModel(req)
	updated.ID = existing.ID

	if err := h.service.Update(updated); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengubah presensi", updated)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Presensi tidak ditemukan")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus presensi")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus presensi", nil)
}
