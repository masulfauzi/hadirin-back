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

// VisibleMenuRow adalah hasil join menus + role_menu_permissions untuk
// satu atau lebih role, dengan permission di-agregasi (OR) lintas role
// memakai bool_or() supaya user dengan banyak role mendapat izin
// gabungan, bukan cuma dari satu role saja.
type VisibleMenuRow struct {
	ID        uuid.UUID
	ParentID  *uuid.UUID
	KodeMenu  string
	NamaMenu  string
	Icon      *string
	Route     *string
	Urutan    int
	CanRead   bool
	CanInsert bool
	CanUpdate bool
	CanDelete bool
}

func (r *Repository) FindVisibleForRoles(roleIDs []uuid.UUID) ([]VisibleMenuRow, error) {
	if len(roleIDs) == 0 {
		return []VisibleMenuRow{}, nil
	}

	var rows []VisibleMenuRow
	err := r.db.Raw(`
		SELECT m.id, m.parent_id, m.kode_menu, m.nama_menu, m.icon, m.route, m.urutan,
		       bool_or(p.can_read) AS can_read,
		       bool_or(p.can_insert) AS can_insert,
		       bool_or(p.can_update) AS can_update,
		       bool_or(p.can_delete) AS can_delete
		FROM menus m
		JOIN role_menu_permissions p ON p.menu_id = m.id
		WHERE m.is_active = true AND p.role_id IN ?
		GROUP BY m.id
		HAVING bool_or(p.can_show) = true
		ORDER BY m.urutan
	`, roleIDs).Scan(&rows).Error
	return rows, err
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
