package harilibur

import (
	"errors"

	"github.com/google/uuid"

	"hadirin-back/modules/division"
)

type Service struct {
	repo            *Repository
	divisionService *division.Service
}

func NewService(repo *Repository, divisionService *division.Service) *Service {
	return &Service{repo: repo, divisionService: divisionService}
}

func (s *Service) GetAll() ([]HariLibur, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*HariLibur, error) {
	return s.repo.FindByID(id)
}

func (s *Service) validate(h *HariLibur) error {
	if h.Tanggal == "" || h.Keterangan == "" {
		return errors.New("tanggal dan keterangan wajib diisi")
	}
	if h.DivisionID != nil {
		if _, err := s.divisionService.GetByID(*h.DivisionID); err != nil {
			return errors.New("division_id tidak ditemukan")
		}
	}
	return nil
}

func (s *Service) Create(h *HariLibur) error {
	if err := s.validate(h); err != nil {
		return err
	}
	return s.repo.Create(h)
}

func (s *Service) Update(h *HariLibur) error {
	if err := s.validate(h); err != nil {
		return err
	}
	return s.repo.Update(h)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
