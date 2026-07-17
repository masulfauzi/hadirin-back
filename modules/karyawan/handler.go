package karyawan

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
	KodeIdentitas  string    `json:"kode_identitas"`
	NIK            *string   `json:"nik"`
	NamaLengkap    string    `json:"nama_lengkap"`
	DivisionID     uuid.UUID `json:"division_id"`
	Jabatan        *string   `json:"jabatan"`
	NoHP           *string   `json:"no_hp"`
	Email          *string   `json:"email"`
	Alamat         *string   `json:"alamat"`
	TanggalMasuk   *string   `json:"tanggal_masuk"`
	StatusKaryawan string    `json:"status_karyawan"`
	FotoProfile    *string   `json:"foto_profile"`
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	list, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data karyawan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data karyawan", list)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	k, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Karyawan tidak ditemukan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil mengambil data karyawan", k)
}

func toModel(req upsertRequest) *Karyawan {
	status := req.StatusKaryawan
	if status == "" {
		status = "aktif"
	}
	return &Karyawan{
		KodeIdentitas:  req.KodeIdentitas,
		NIK:            req.NIK,
		NamaLengkap:    req.NamaLengkap,
		DivisionID:     req.DivisionID,
		Jabatan:        req.Jabatan,
		NoHP:           req.NoHP,
		Email:          req.Email,
		Alamat:         req.Alamat,
		TanggalMasuk:   req.TanggalMasuk,
		StatusKaryawan: status,
		FotoProfile:    req.FotoProfile,
	}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req upsertRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	k := toModel(req)
	if err := h.service.Create(k); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.Success(c, fiber.StatusCreated, "Berhasil membuat karyawan", k)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	existing, err := h.service.GetByID(id)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Karyawan tidak ditemukan")
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
	return utils.Success(c, fiber.StatusOK, "Berhasil mengubah karyawan", updated)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if _, err := h.service.GetByID(id); err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Karyawan tidak ditemukan")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal menghapus karyawan")
	}
	return utils.Success(c, fiber.StatusOK, "Berhasil menghapus karyawan", nil)
}
