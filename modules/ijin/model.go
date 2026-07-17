package ijin

import (
	"time"

	"github.com/google/uuid"
)

type Ijin struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KaryawanID       uuid.UUID  `gorm:"type:uuid;not null" json:"karyawan_id"`
	JenisIjinID      uuid.UUID  `gorm:"type:uuid;not null" json:"jenis_ijin_id"`
	StatusIjinID     uuid.UUID  `gorm:"type:uuid;not null" json:"status_ijin_id"`
	TanggalMulai     string     `gorm:"not null" json:"tanggal_mulai"`
	TanggalSelesai   string     `gorm:"not null" json:"tanggal_selesai"`
	JumlahHari       int        `gorm:"not null;default:1" json:"jumlah_hari"`
	Alasan           *string    `json:"alasan"`
	FileLampiran     *string    `gorm:"size:255" json:"file_lampiran"`
	DisetujuiOleh    *uuid.UUID `gorm:"type:uuid" json:"disetujui_oleh"`
	TanggalApproval  *time.Time `json:"tanggal_approval"`
	CatatanApproval  *string    `gorm:"size:255" json:"catatan_approval"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

func (Ijin) TableName() string { return "ijin" }

type StatusIjin struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KodeStatus  string    `gorm:"size:20;uniqueIndex;not null" json:"kode_status"`
	NamaStatus  string    `gorm:"size:50;not null" json:"nama_status"`
	WarnaBadge  *string   `gorm:"size:20" json:"warna_badge"`
	Urutan      int       `gorm:"not null;default:0" json:"urutan"`
}

func (StatusIjin) TableName() string { return "status_ijin" }

type JenisIjin struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	KodeJenis      string    `gorm:"size:20;uniqueIndex;not null" json:"kode_jenis"`
	NamaJenis      string    `gorm:"size:100;not null" json:"nama_jenis"`
	MaxHari        *int      `json:"max_hari"`
	PerluLampiran  bool      `gorm:"not null;default:false" json:"perlu_lampiran"`
}

func (JenisIjin) TableName() string { return "jenis_ijin" }
