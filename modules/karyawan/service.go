package karyawan

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

func (s *Service) GetAll() ([]Karyawan, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Karyawan, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByKodeIdentitas(kodeIdentitas string) (*Karyawan, error) {
	return s.repo.FindByKodeIdentitas(kodeIdentitas)
}

// CreateFromRegistration membuat baris karyawan saat user mendaftar lewat
// POST /auth/register, dengan division_id NULL karena divisi belum
// ditentukan admin. Sengaja tidak lewat validate() karena division_id
// wajib diisi untuk endpoint /karyawan biasa, tapi tidak di alur ini.
func (s *Service) CreateFromRegistration(kodeIdentitas, namaLengkap string) (*Karyawan, error) {
	k := &Karyawan{
		KodeIdentitas:  kodeIdentitas,
		NamaLengkap:    namaLengkap,
		StatusKaryawan: "aktif",
	}
	if err := s.repo.Create(k); err != nil {
		return nil, err
	}
	return k, nil
}

func (s *Service) validate(k *Karyawan) error {
	if k.KodeIdentitas == "" || k.NamaLengkap == "" {
		return errors.New("kode_identitas dan nama_lengkap wajib diisi")
	}
	if k.DivisionID == nil {
		return errors.New("division_id wajib diisi")
	}
	if _, err := s.divisionService.GetByID(*k.DivisionID); err != nil {
		return errors.New("division_id tidak ditemukan")
	}
	return nil
}

func (s *Service) Create(k *Karyawan) error {
	if err := s.validate(k); err != nil {
		return err
	}
	return s.repo.Create(k)
}

func (s *Service) Update(k *Karyawan) error {
	if err := s.validate(k); err != nil {
		return err
	}
	return s.repo.Update(k)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
