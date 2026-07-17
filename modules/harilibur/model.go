package harilibur

import (
	"time"

	"github.com/google/uuid"
)

type HariLibur struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DivisionID  *uuid.UUID `gorm:"type:uuid" json:"division_id"`
	Tanggal     string     `gorm:"not null" json:"tanggal"`
	Keterangan  string     `gorm:"size:150;not null" json:"keterangan"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (HariLibur) TableName() string { return "hari_libur" }
