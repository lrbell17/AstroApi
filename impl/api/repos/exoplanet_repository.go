package repos

import (
	"github.com/lrbell17/astroapi/impl/model"
	"gorm.io/gorm"
)

type ExoplanetRepo struct {
	db *gorm.DB
}

func NewExoplanetRepo(db *gorm.DB) *ExoplanetRepo {
	return &ExoplanetRepo{db}
}

// Get planet by ID
func (r *ExoplanetRepo) GetById(id uint) (*model.Exoplanet, error) {
	var planet model.Exoplanet
	result := r.db.Preload("Star").First(&planet, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &planet, nil
}
