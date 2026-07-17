package ijin

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
	KaryawanID      uuid.UUID `json:"karyawan_id"`
	JenisIjinID     uuid.UUID `json:"jenis_ijin_id"`
	StatusIjinID    uuid.UUID `json:"status_ijin_id"`
	TanggalMulai    string    `json:"tanggal_mulai"`
	TanggalSelesai  string    `json:"tanggal_selesai"`
	JumlahHari      int       `json:"jumlah_hari"`
	Alasan          *string   `json:"alasan"`
	FileLampiran    *string   `json:"file_lampiran"`
}

func toModel(req upsertRequest) *Ijin {
	jumlahHari := req.JumlahHari
	if jumlahHari <= 0 {
		jumlahHari = 1
	}
	return &Ijin{
		KaryawanID:     req.KaryawanID,
		JenisIjinID:    req.JenisIjinID,
		StatusIjinID:   req.StatusIjinID,
		TanggalMulai:   req.TanggalMulai,
		TanggalSelesai: req.TanggalSelesai,
		JumlahHari:     jumlahHari,
		Alasan:         req.Alasan,
		FileLampiran:   req.FileLampiran,
	}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	list, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data ijin")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data ijin", list)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	i, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Ijin tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data ijin", i)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	i := toModel(req)
	if err := h.service.Create(i); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil membuat ijin", i)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	existing, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Ijin tidak ditemukan")
	}

	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	updated := toModel(req)
	updated.ID = existing.ID
	updated.DisetujuiOleh = existing.DisetujuiOleh
	updated.TanggalApproval = existing.TanggalApproval
	updated.CatatanApproval = existing.CatatanApproval

	if err := h.service.Update(updated); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengubah ijin", updated)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Ijin tidak ditemukan")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus ijin")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus ijin", nil)
}

func (h *Handler) GetStatusList(c *fiber.Ctx) error {
	list, err := h.service.GetAllStatus()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data status ijin")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data status ijin", list)
}

func (h *Handler) GetJenisList(c *fiber.Ctx) error {
	list, err := h.service.GetAllJenis()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data jenis ijin")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data jenis ijin", list)
}
