package division

import (
	"time"

	"github.com/google/uuid"
)

type Division struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KodeDivisi  string    `gorm:"size:50;uniqueIndex;not null" json:"kode_divisi"`
	NamaDivisi  string    `gorm:"size:100;not null" json:"nama_divisi"`
	Deskripsi   *string   `gorm:"size:255" json:"deskripsi"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Division) TableName() string { return "divisions" }

type Schedule struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DivisionID      uuid.UUID `gorm:"type:uuid;not null" json:"division_id"`
	Hari            string    `gorm:"size:10;not null" json:"hari"`
	IsHariKerja     bool      `gorm:"not null;default:true" json:"is_hari_kerja"`
	JamMasuk        *string   `json:"jam_masuk"`
	JamKeluar       *string   `json:"jam_keluar"`
	ToleransiMenit  int       `gorm:"not null;default:0" json:"toleransi_menit"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (Schedule) TableName() string { return "division_schedules" }
