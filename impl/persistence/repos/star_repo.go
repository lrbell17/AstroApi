package repos

import (
	"github.com/lrbell17/astroapi/impl/persistence/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StarRepo struct {
	db *gorm.DB
}

// Constructor for star repo
func NewStarRepo(db *gorm.DB) *StarRepo {
	return &StarRepo{db}
}

// Get star by ID from DB
func (r *StarRepo) GetById(id uint) (*dao.Star, error) {
	var star dao.Star
	result := r.db.Preload("Exoplanets").First(&star, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &star, nil
}

// Add star to DB
func (r *StarRepo) Insert(s *dao.Star) (*dao.Star, error) {

	if err := r.db.Create(s).Error; err != nil {
		log.Errorf("Error adding star: %v", err)
		return nil, err
	}

	return s, nil
}

// Insert batch of stars
func (r *StarRepo) BatchInsert(stars []*dao.Star) (int, error) {

	result := r.db.Model(&dao.Star{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "name"}}, DoNothing: true}). // Ignore duplicates
		CreateInBatches(stars, len(stars))

	return int(result.RowsAffected), result.Error
}

// Get all stars
func (r *StarRepo) GetAll() (stars []dao.Star, err error) {
	result := r.db.Preload("Exoplanets").Find(&stars)
	err = result.Error
	return
}

// Search for stars by name with partial matching
func (r *StarRepo) SearchByName(name string, limit int) (stars []dao.Star, err error) {
	result := r.db.Where("name ILIKE ?", "%"+name+"%").Limit(limit).Find(&stars)
	err = result.Error
	return
}
