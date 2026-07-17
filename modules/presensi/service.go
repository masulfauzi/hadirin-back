package presensi

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

func (s *Service) GetAll() ([]Presensi, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Presensi, error) {
	return s.repo.FindByID(id)
}

func (s *Service) validate(p *Presensi) error {
	if p.Tanggal == "" {
		return errors.New("tanggal wajib diisi")
	}
	if _, err := s.karyawanService.GetByID(p.KaryawanID); err != nil {
		return errors.New("karyawan_id tidak ditemukan")
	}
	return nil
}

func (s *Service) Create(p *Presensi) error {
	if err := s.validate(p); err != nil {
		return err
	}
	return s.repo.Create(p)
}

func (s *Service) Update(p *Presensi) error {
	if err := s.validate(p); err != nil {
		return err
	}
	return s.repo.Update(p)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
