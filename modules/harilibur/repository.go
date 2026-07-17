package harilibur

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

func (r *Repository) FindAll() ([]HariLibur, error) {
	var list []HariLibur
	err := r.db.Find(&list).Error
	return list, err
}

func (r *Repository) FindByID(id uuid.UUID) (*HariLibur, error) {
	var h HariLibur
	if err := r.db.First(&h, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *Repository) Create(h *HariLibur) error {
	return r.db.Create(h).Error
}

func (r *Repository) Update(h *HariLibur) error {
	return r.db.Save(h).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&HariLibur{}, "id = ?", id).Error
}
