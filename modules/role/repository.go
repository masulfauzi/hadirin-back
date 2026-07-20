package role

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Role, error) {
	var roles []Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *Repository) FindByID(id uuid.UUID) (*Role, error) {
	var role Role
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Repository) Create(role *Role) error {
	return r.db.Create(role).Error
}

func (r *Repository) Update(role *Role) error {
	return r.db.Save(role).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Role{}, "id = ?", id).Error
}

func (r *Repository) AssignUser(roleID, userID uuid.UUID) error {
	ur := UserRole{RoleID: roleID, UserID: userID}
	return r.db.Create(&ur).Error
}

func (r *Repository) RevokeUser(roleID, userID uuid.UUID) error {
	return r.db.Where("role_id = ? AND user_id = ?", roleID, userID).Delete(&UserRole{}).Error
}

func (r *Repository) FindByKodeRole(kodeRole string) (*Role, error) {
	var role Role
	if err := r.db.Where("kode_role = ?", kodeRole).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Repository) FindRoleIDsByUserID(userID uuid.UUID) ([]uuid.UUID, error) {
	var roleIDs []uuid.UUID
	err := r.db.Model(&UserRole{}).Where("user_id = ?", userID).Pluck("role_id", &roleIDs).Error
	return roleIDs, err
}
