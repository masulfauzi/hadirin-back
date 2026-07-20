package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"hadirin-back/config"
	"hadirin-back/modules/karyawan"
	"hadirin-back/modules/role"
	"hadirin-back/modules/user"
)

// kodeRoleKaryawan adalah role default yang di-assign ke setiap user baru
// lewat /auth/register (di-seed lewat migration 000017_seed_karyawan_role).
const kodeRoleKaryawan = "KARYAWAN"

type Service struct {
	userService     *user.Service
	karyawanService *karyawan.Service
	roleService     *role.Service
	cfg             *config.Config
}

func NewService(userService *user.Service, karyawanService *karyawan.Service, roleService *role.Service, cfg *config.Config) *Service {
	return &Service{userService: userService, karyawanService: karyawanService, roleService: roleService, cfg: cfg}
}

func (s *Service) Register(kodeIdentitas, username string, email *string, password, namaLengkap string) (*user.User, error) {
	if _, err := s.userService.GetUserByUsername(username); err == nil {
		return nil, errors.New("username sudah terdaftar")
	}
	if _, err := s.userService.GetUserByKodeIdentitas(kodeIdentitas); err == nil {
		return nil, errors.New("kode_identitas sudah terdaftar")
	}
	if _, err := s.karyawanService.GetByKodeIdentitas(kodeIdentitas); err == nil {
		return nil, errors.New("kode_identitas sudah terdaftar sebagai karyawan")
	}

	karyawanRole, err := s.roleService.GetByKodeRole(kodeRoleKaryawan)
	if err != nil {
		return nil, errors.New("role karyawan belum tersedia, hubungi admin")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		KodeIdentitas: kodeIdentitas,
		Username:      username,
		Email:         email,
		Password:      string(hashed),
	}
	if err := s.userService.CreateUser(newUser); err != nil {
		return nil, err
	}

	newKaryawan, err := s.karyawanService.CreateFromRegistration(kodeIdentitas, namaLengkap)
	if err != nil {
		// Kompensasi: rollback pembuatan user kalau baris karyawan gagal
		// dibuat, supaya tidak ada user tanpa data karyawan pasangannya.
		_ = s.userService.DeleteUser(newUser.ID)
		return nil, err
	}

	if err := s.roleService.AssignUser(karyawanRole.ID, newUser.ID); err != nil {
		// Kompensasi: rollback user + karyawan kalau assign role gagal,
		// supaya tidak ada user yang lolos registrasi tanpa role.
		_ = s.karyawanService.Delete(newKaryawan.ID)
		_ = s.userService.DeleteUser(newUser.ID)
		return nil, err
	}

	return newUser, nil
}

func (s *Service) Login(username, password string) (string, *user.User, error) {
	// Pesan error sengaja disamakan untuk username tak terdaftar maupun
	// password salah, agar penyerang tidak bisa menebak username mana
	// yang terdaftar
	loggedUser, err := s.userService.GetUserByUsername(username)
	if err != nil {
		return "", nil, errors.New("username atau password salah")
	}
	if bcrypt.CompareHashAndPassword([]byte(loggedUser.Password), []byte(password)) != nil {
		return "", nil, errors.New("username atau password salah")
	}
	if !loggedUser.IsActive {
		return "", nil, errors.New("akun tidak aktif")
	}

	claims := jwt.MapClaims{
		"user_id": loggedUser.ID.String(),
		"exp":     time.Now().Add(time.Duration(s.cfg.JWTExpireHours) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	if err := s.userService.UpdateLastLogin(loggedUser); err != nil {
		return "", nil, err
	}

	return signed, loggedUser, nil
}

func (s *Service) GetProfile(id uuid.UUID) (*user.User, error) {
	return s.userService.GetUserByID(id)
}
