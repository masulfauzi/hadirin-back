package user

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *Service) GetUserByID(id uuid.UUID) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetUserByUsername(username string) (*User, error) {
	return s.repo.FindByUsername(username)
}

func (s *Service) GetUserByKodeIdentitas(kodeIdentitas string) (*User, error) {
	return s.repo.FindByKodeIdentitas(kodeIdentitas)
}

func (s *Service) CreateUser(u *User) error {
	return s.repo.Create(u)
}

func (s *Service) UpdateUser(u *User) error {
	return s.repo.Update(u)
}

func (s *Service) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *Service) UpdateLastLogin(u *User) error {
	now := time.Now()
	u.LastLoginAt = &now
	return s.repo.Update(u)
}
