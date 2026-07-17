package role

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KodeRole   string    `gorm:"size:50;uniqueIndex;not null" json:"kode_role"`
	NamaRole   string    `gorm:"size:100;not null" json:"nama_role"`
	Deskripsi  *string   `gorm:"size:255" json:"deskripsi"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Role) TableName() string { return "roles" }

type UserRole struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	RoleID    uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (UserRole) TableName() string { return "user_roles" }
