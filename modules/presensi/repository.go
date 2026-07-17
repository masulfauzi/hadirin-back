package presensi

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

func (r *Repository) FindAll() ([]Presensi, error) {
	var list []Presensi
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) FindByID(id uuid.UUID) (*Presensi, error) {
	var p Presensi
	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Create(p *Presensi) error {
	return r.db.Create(p).Error
}

func (r *Repository) Update(p *Presensi) error {
	return r.db.Save(p).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Presensi{}, "id = ?", id).Error
}
