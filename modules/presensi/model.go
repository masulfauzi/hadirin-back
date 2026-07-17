package presensi

import (
	"time"

	"github.com/google/uuid"
)

type Presensi struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KaryawanID  uuid.UUID `gorm:"type:uuid;not null" json:"karyawan_id"`
	Tanggal     string    `gorm:"not null" json:"tanggal"`

	JamHadir         *time.Time `json:"jam_hadir"`
	FotoHadir        *string    `gorm:"size:255" json:"foto_hadir"`
	LatHadir         *float64   `json:"lat_hadir"`
	LongHadir        *float64   `json:"long_hadir"`
	JarakHadirMeter  *float64   `json:"jarak_hadir_meter"`
	StatusHadir      *string    `gorm:"size:20" json:"status_hadir"`

	JamPulang         *time.Time `json:"jam_pulang"`
	FotoPulang        *string    `gorm:"size:255" json:"foto_pulang"`
	LatPulang         *float64   `json:"lat_pulang"`
	LongPulang        *float64   `json:"long_pulang"`
	JarakPulangMeter  *float64   `json:"jarak_pulang_meter"`
	StatusPulang      *string    `gorm:"size:20" json:"status_pulang"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Presensi) TableName() string { return "presensi" }
