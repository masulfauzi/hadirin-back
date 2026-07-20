package role

import (
	"errors"

	"github.com/google/uuid"

	"hadirin-back/modules/user"
)

type Service struct {
	repo        *Repository
	userService *user.Service
}

func NewService(repo *Repository, userService *user.Service) *Service {
	return &Service{repo: repo, userService: userService}
}

func (s *Service) GetAll() ([]Role, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Role, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(role *Role) error {
	if role.KodeRole == "" || role.NamaRole == "" {
		return errors.New("kode_role dan nama_role wajib diisi")
	}
	return s.repo.Create(role)
}

func (s *Service) Update(role *Role) error {
	if role.KodeRole == "" || role.NamaRole == "" {
		return errors.New("kode_role dan nama_role wajib diisi")
	}
	return s.repo.Update(role)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *Service) AssignUser(roleID, userID uuid.UUID) error {
	if _, err := s.repo.FindByID(roleID); err != nil {
		return errors.New("role tidak ditemukan")
	}
	if _, err := s.userService.GetUserByID(userID); err != nil {
		return errors.New("user tidak ditemukan")
	}
	return s.repo.AssignUser(roleID, userID)
}

func (s *Service) RevokeUser(roleID, userID uuid.UUID) error {
	return s.repo.RevokeUser(roleID, userID)
}

func (s *Service) GetByKodeRole(kodeRole string) (*Role, error) {
	return s.repo.FindByKodeRole(kodeRole)
}

func (s *Service) GetRoleIDsByUserID(userID uuid.UUID) ([]uuid.UUID, error) {
	return s.repo.FindRoleIDsByUserID(userID)
}
