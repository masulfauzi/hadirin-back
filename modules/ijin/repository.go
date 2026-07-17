package ijin

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

func (r *Repository) FindAll() ([]Ijin, error) {
	var list []Ijin
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) FindByID(id uuid.UUID) (*Ijin, error) {
	var i Ijin
	if err := r.db.First(&i, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &i, nil
}

func (r *Repository) Create(i *Ijin) error {
	return r.db.Create(i).Error
}

func (r *Repository) Update(i *Ijin) error {
	return r.db.Save(i).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Ijin{}, "id = ?", id).Error
}

func (r *Repository) FindByJenisID(id uuid.UUID) (*JenisIjin, error) {
	var j JenisIjin
	if err := r.db.First(&j, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &j, nil
}

func (r *Repository) FindByStatusID(id uuid.UUID) (*StatusIjin, error) {
	var s StatusIjin
	if err := r.db.First(&s, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *Repository) FindAllStatus() ([]StatusIjin, error) {
	var list []StatusIjin
	err := r.db.Order("urutan").Find(&list).Error
	return list, err
}

func (r *Repository) FindAllJenis() ([]JenisIjin, error) {
	var list []JenisIjin
	err := r.db.Find(&list).Error
	return list, err
}
