package karyawan

import (
	"time"

	"github.com/google/uuid"
)

type Karyawan struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KodeIdentitas  string     `gorm:"size:50;uniqueIndex;not null" json:"kode_identitas"`
	NamaLengkap    string     `gorm:"size:150;not null" json:"nama_lengkap"`
	DivisionID     *uuid.UUID `gorm:"type:uuid" json:"division_id"`
	Jabatan        *string    `gorm:"size:100" json:"jabatan"`
	NoHP           *string    `gorm:"column:no_hp;size:20" json:"no_hp"`
	Email          *string    `gorm:"size:100" json:"email"`
	Alamat         *string    `json:"alamat"`
	TanggalMasuk   *string    `json:"tanggal_masuk"`
	StatusKaryawan string     `gorm:"size:20;not null;default:aktif" json:"status_karyawan"`
	FotoProfile    *string    `gorm:"size:255" json:"foto_profile"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (Karyawan) TableName() string { return "karyawan" }
