package menu

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

func (r *Repository) FindAll() ([]Menu, error) {
	var menus []Menu
	err := r.db.Find(&menus).Error
	return menus, err
}

func (r *Repository) FindByID(id uuid.UUID) (*Menu, error) {
	var m Menu
	if err := r.db.First(&m, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *Repository) Create(m *Menu) error {
	return r.db.Create(m).Error
}

func (r *Repository) Update(m *Menu) error {
	return r.db.Save(m).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Menu{}, "id = ?", id).Error
}

func (r *Repository) FindPermissionsByMenuID(menuID uuid.UUID) ([]RoleMenuPermission, error) {
	var perms []RoleMenuPermission
	err := r.db.Where("menu_id = ?", menuID).Find(&perms).Error
	return perms, err
}

func (r *Repository) UpsertPermission(p *RoleMenuPermission) error {
	var existing RoleMenuPermission
	err := r.db.Where("role_id = ? AND menu_id = ?", p.RoleID, p.MenuID).First(&existing).Error
	if err == nil {
		p.ID = existing.ID
		return r.db.Save(p).Error
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	return r.db.Create(p).Error
}
