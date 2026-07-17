package division

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"hadirin-back/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type upsertRequest struct {
	KodeDivisi string  `json:"kode_divisi"`
	NamaDivisi string  `json:"nama_divisi"`
	Deskripsi  *string `json:"deskripsi"`
	IsActive   *bool   `json:"is_active"`
}

type scheduleRequest struct {
	Hari            string  `json:"hari"`
	IsHariKerja     bool    `json:"is_hari_kerja"`
	JamMasuk        *string `json:"jam_masuk"`
	JamKeluar       *string `json:"jam_keluar"`
	ToleransiMenit  int     `json:"toleransi_menit"`
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	divisions, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data divisi")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data divisi", divisions)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	d, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Divisi tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data divisi", d)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	d := &Division{
		KodeDivisi: req.KodeDivisi,
		NamaDivisi: req.NamaDivisi,
		Deskripsi:  req.Deskripsi,
	}
	if req.IsActive != nil {
		d.IsActive = *req.IsActive
	} else {
		d.IsActive = true
	}

	if err := h.service.Create(d); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil membuat divisi", d)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	d, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Divisi tidak ditemukan")
	}

	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	d.KodeDivisi = req.KodeDivisi
	d.NamaDivisi = req.NamaDivisi
	d.Deskripsi = req.Deskripsi
	if req.IsActive != nil {
		d.IsActive = *req.IsActive
	}

	if err := h.service.Update(d); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengubah divisi", d)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Divisi tidak ditemukan")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus divisi")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus divisi", nil)
}

func (h *Handler) GetSchedules(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	schedules, err := h.service.GetSchedules(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Error(c, fiber.StatusNotFound, "Divisi tidak ditemukan")
		}
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil jadwal divisi")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil jadwal divisi", schedules)
}

func (h *Handler) SetSchedules(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	var reqs []scheduleRequest
	if err := c.BodyParser(&reqs); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	schedules := make([]Schedule, len(reqs))
	for i, r := range reqs {
		schedules[i] = Schedule{
			Hari:           r.Hari,
			IsHariKerja:    r.IsHariKerja,
			JamMasuk:       r.JamMasuk,
			JamKeluar:      r.JamKeluar,
			ToleransiMenit: r.ToleransiMenit,
		}
	}

	if err := h.service.SetSchedules(id, schedules); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.Error(c, fiber.StatusNotFound, "Divisi tidak ditemukan")
		}
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}

	updated, err := h.service.GetSchedules(id)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil jadwal divisi")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menyimpan jadwal divisi", updated)
}
