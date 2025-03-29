package repos

import (
	"github.com/lrbell17/astroapi/impl/persistence/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ExoplanetRepo struct {
	db *gorm.DB
}

// Constructor for exoplanet repo
func NewExoplanetRepo(db *gorm.DB) *ExoplanetRepo {
	return &ExoplanetRepo{db}
}

// Get planet by ID from database
func (r *ExoplanetRepo) GetById(id uint) (*dao.Exoplanet, error) {
	var planet dao.Exoplanet
	result := r.db.Preload("Star").First(&planet, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &planet, nil
}

// Add planet to repo
func (r *ExoplanetRepo) Insert(e *dao.Exoplanet) (*dao.Exoplanet, error) {

	if err := r.db.Create(e).Error; err != nil {
		log.Errorf("Error adding planet: %v", err)
		return nil, err
	}

	return e, nil
}

// Insert batch of exoplanets
func (r *ExoplanetRepo) BatchInsert(planets []*dao.Exoplanet) (int, error) {
	result := r.db.Model(&dao.Exoplanet{}).CreateInBatches(planets, len(planets))
	return int(result.RowsAffected), result.Error
}
