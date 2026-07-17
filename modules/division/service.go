package division

import (
	"errors"

	"github.com/google/uuid"
)

var validHari = map[string]bool{
	"senin": true, "selasa": true, "rabu": true, "kamis": true,
	"jumat": true, "sabtu": true, "minggu": true,
}

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]Division, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Division, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(d *Division) error {
	if d.KodeDivisi == "" || d.NamaDivisi == "" {
		return errors.New("kode_divisi dan nama_divisi wajib diisi")
	}
	return s.repo.Create(d)
}

func (s *Service) Update(d *Division) error {
	if d.KodeDivisi == "" || d.NamaDivisi == "" {
		return errors.New("kode_divisi dan nama_divisi wajib diisi")
	}
	return s.repo.Update(d)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *Service) GetSchedules(divisionID uuid.UUID) ([]Schedule, error) {
	if _, err := s.repo.FindByID(divisionID); err != nil {
		return nil, err
	}
	return s.repo.FindSchedulesByDivisionID(divisionID)
}

func (s *Service) SetSchedules(divisionID uuid.UUID, schedules []Schedule) error {
	if _, err := s.repo.FindByID(divisionID); err != nil {
		return err
	}
	for i := range schedules {
		if !validHari[schedules[i].Hari] {
			return errors.New("hari tidak valid: " + schedules[i].Hari)
		}
		schedules[i].DivisionID = divisionID
	}
	return s.repo.ReplaceSchedules(divisionID, schedules)
}
