package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"hadirin-back/config"
	"hadirin-back/modules/user"
)

type Service struct {
	userService *user.Service
	cfg         *config.Config
}

func NewService(userService *user.Service, cfg *config.Config) *Service {
	return &Service{userService: userService, cfg: cfg}
}

func (s *Service) Register(name, email, password string) (*user.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}
	if err := s.userService.CreateUser(newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *Service) Login(email, password string) (string, *user.User, error) {
	// Pesan error sengaja disamakan untuk email tak terdaftar maupun
	// password salah, agar penyerang tidak bisa menebak email mana
	// yang terdaftar
	loggedUser, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return "", nil, errors.New("email atau password salah")
	}
	if bcrypt.CompareHashAndPassword([]byte(loggedUser.Password), []byte(password)) != nil {
		return "", nil, errors.New("email atau password salah")
	}

	claims := jwt.MapClaims{
		"user_id": loggedUser.ID,
		"exp":     time.Now().Add(time.Duration(s.cfg.JWTExpireHours) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	return signed, loggedUser, nil
}

func (s *Service) GetProfile(id uint) (*user.User, error) {
	return s.userService.GetUserByID(id)
}
