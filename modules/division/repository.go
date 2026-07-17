package division

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

func (r *Repository) FindAll() ([]Division, error) {
	var divisions []Division
	err := r.db.Find(&divisions).Error
	return divisions, err
}

func (r *Repository) FindByID(id uuid.UUID) (*Division, error) {
	var d Division
	if err := r.db.First(&d, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *Repository) Create(d *Division) error {
	return r.db.Create(d).Error
}

func (r *Repository) Update(d *Division) error {
	return r.db.Save(d).Error
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Division{}, "id = ?", id).Error
}

func (r *Repository) FindSchedulesByDivisionID(divisionID uuid.UUID) ([]Schedule, error) {
	var schedules []Schedule
	err := r.db.Where("division_id = ?", divisionID).Find(&schedules).Error
	return schedules, err
}

// ReplaceSchedules menghapus jadwal lama divisi lalu menyimpan jadwal baru
// dalam satu transaksi, dipakai untuk endpoint set jadwal 7 hari sekaligus.
func (r *Repository) ReplaceSchedules(divisionID uuid.UUID, schedules []Schedule) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("division_id = ?", divisionID).Delete(&Schedule{}).Error; err != nil {
			return err
		}
		if len(schedules) == 0 {
			return nil
		}
		return tx.Create(&schedules).Error
	})
}
