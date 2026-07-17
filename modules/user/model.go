package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KodeIdentitas string     `gorm:"size:50;uniqueIndex;not null" json:"kode_identitas"`
	Username      string     `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email         *string    `gorm:"size:100;uniqueIndex" json:"email"`
	Password      string     `gorm:"column:password_hash;size:255;not null" json:"-"`
	IsActive      bool       `gorm:"not null;default:true" json:"is_active"`
	LastLoginAt   *time.Time `json:"last_login_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
