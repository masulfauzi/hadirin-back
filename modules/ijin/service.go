package ijin

import (
	"errors"

	"github.com/google/uuid"

	"hadirin-back/modules/karyawan"
)

type Service struct {
	repo            *Repository
	karyawanService *karyawan.Service
}

func NewService(repo *Repository, karyawanService *karyawan.Service) *Service {
	return &Service{repo: repo, karyawanService: karyawanService}
}

func (s *Service) GetAll() ([]Ijin, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Ijin, error) {
	return s.repo.FindByID(id)
}

func (s *Service) validate(i *Ijin) error {
	if i.TanggalMulai == "" || i.TanggalSelesai == "" {
		return errors.New("tanggal_mulai dan tanggal_selesai wajib diisi")
	}
	if i.TanggalSelesai < i.TanggalMulai {
		return errors.New("tanggal_selesai tidak boleh sebelum tanggal_mulai")
	}
	if _, err := s.karyawanService.GetByID(i.KaryawanID); err != nil {
		return errors.New("karyawan_id tidak ditemukan")
	}
	if _, err := s.repo.FindByJenisID(i.JenisIjinID); err != nil {
		return errors.New("jenis_ijin_id tidak ditemukan")
	}
	if _, err := s.repo.FindByStatusID(i.StatusIjinID); err != nil {
		return errors.New("status_ijin_id tidak ditemukan")
	}
	return nil
}

func (s *Service) Create(i *Ijin) error {
	if err := s.validate(i); err != nil {
		return err
	}
	return s.repo.Create(i)
}

func (s *Service) Update(i *Ijin) error {
	if err := s.validate(i); err != nil {
		return err
	}
	return s.repo.Update(i)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *Service) GetAllStatus() ([]StatusIjin, error) {
	return s.repo.FindAllStatus()
}

func (s *Service) GetAllJenis() ([]JenisIjin, error) {
	return s.repo.FindAllJenis()
}
