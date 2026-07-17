package menu

import (
	"errors"

	"github.com/google/uuid"

	"hadirin-back/modules/role"
)

type Service struct {
	repo        *Repository
	roleService *role.Service
}

func NewService(repo *Repository, roleService *role.Service) *Service {
	return &Service{repo: repo, roleService: roleService}
}

func (s *Service) GetAll() ([]Menu, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Menu, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(m *Menu) error {
	if m.KodeMenu == "" || m.NamaMenu == "" {
		return errors.New("kode_menu dan nama_menu wajib diisi")
	}
	return s.repo.Create(m)
}

func (s *Service) Update(m *Menu) error {
	if m.KodeMenu == "" || m.NamaMenu == "" {
		return errors.New("kode_menu dan nama_menu wajib diisi")
	}
	return s.repo.Update(m)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *Service) GetPermissions(menuID uuid.UUID) ([]RoleMenuPermission, error) {
	if _, err := s.repo.FindByID(menuID); err != nil {
		return nil, err
	}
	return s.repo.FindPermissionsByMenuID(menuID)
}

func (s *Service) SetPermission(menuID uuid.UUID, p *RoleMenuPermission) error {
	if _, err := s.repo.FindByID(menuID); err != nil {
		return errors.New("menu tidak ditemukan")
	}
	if _, err := s.roleService.GetByID(p.RoleID); err != nil {
		return errors.New("role_id tidak ditemukan")
	}
	p.MenuID = menuID
	return s.repo.UpsertPermission(p)
}
