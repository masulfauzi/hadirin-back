package karyawan

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

func (r *Repository) FindAll() ([]Karyawan, error) {
	var list []Karyawan
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) FindByID(id uuid.UUID) (*Karyawan, error) {
	var k Karyawan
	if err := r.db.First(&k, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *Repository) FindByKodeIdentitas(kodeIdentitas string) (*Karyawan, error) {
	var k Karyawan
	if err := r.db.Where("kode_identitas = ?", kodeIdentitas).First(&k).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *Repository) Create(k *Karyawan) error {
	return r.db.Create(k).Error
}

func (r *Repository) Update(k *Karyawan) error {
	return r.db.Save(k).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Karyawan{}, "id = ?", id).Error
}
