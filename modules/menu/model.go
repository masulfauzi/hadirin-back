package menu

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ParentID  *uuid.UUID `gorm:"type:uuid" json:"parent_id"`
	KodeMenu  string     `gorm:"size:50;uniqueIndex;not null" json:"kode_menu"`
	NamaMenu  string     `gorm:"size:100;not null" json:"nama_menu"`
	Icon      *string    `gorm:"size:100" json:"icon"`
	Route     *string    `gorm:"size:255" json:"route"`
	Urutan    int        `gorm:"not null;default:0" json:"urutan"`
	IsActive  bool       `gorm:"not null;default:true" json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (Menu) TableName() string { return "menus" }

type RoleMenuPermission struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	RoleID     uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	MenuID     uuid.UUID `gorm:"type:uuid;not null" json:"menu_id"`
	CanShow    bool      `gorm:"not null;default:false" json:"can_show"`
	CanRead    bool      `gorm:"not null;default:false" json:"can_read"`
	CanInsert  bool      `gorm:"not null;default:false" json:"can_insert"`
	CanUpdate  bool      `gorm:"not null;default:false" json:"can_update"`
	CanDelete  bool      `gorm:"not null;default:false" json:"can_delete"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (RoleMenuPermission) TableName() string { return "role_menu_permissions" }
