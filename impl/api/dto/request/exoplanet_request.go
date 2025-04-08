package request

import (
	"github.com/lrbell17/astroapi/impl/persistence/dao"
)

type (
	// Exoplanet request interface
	ExoplanetRequest interface {
		Request[dao.Exoplanet]
	}
	// Exoplanet Request DTO
	ExoplanetRequestDTO struct {
		Name          string  `json:"name" binding:"required"`
		StarId        uint    `json:"star_id" binding:"required"`
		Mass          float32 `json:"mass" binding:"gt=0"`
		Radius        float32 `json:"radius" binding:"gt=0"`
		Dist          float32 `json:"dist" binding:"gt=0"`
		OrbitalPeriod float32 `json:"orbital_period" binding:"gt=0"`
	}
)

// Get Exoplanet DAO from Exoplanet request DTO
func (req *ExoplanetRequestDTO) DaoFromRequest() *dao.Exoplanet {
	if req == nil {
		return nil
	}

	return &dao.Exoplanet{
		Name:          req.Name,
		StarID:        req.StarId,
		Mass:          req.Mass,
		Radius:        req.Radius,
		Dist:          req.Dist,
		OrbitalPeriod: req.OrbitalPeriod,
	}
}
