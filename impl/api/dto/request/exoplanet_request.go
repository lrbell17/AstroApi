package request

import (
	"github.com/lrbell17/astroapi/impl/model"
)

type (
	// Exoplanet request interface
	ExoplanetRequest interface {
		Request[model.Exoplanet]
	}
	// Exoplanet Request DTO
	ExoplanetRequestDTO struct {
		Name   string  `json:"name" binding:"required"`
		StarId uint    `json:"star_id" binding:"required"`
		Mass   float32 `json:"mass" binding:"gt=0"`
		Radius float32 `json:"radius" binding:"gt=0"`
		Dist   float32 `json:"dist" binding:"gt=0"`
	}
)

// Get Exoplanet model from Exoplanet request DTO
func (req *ExoplanetRequestDTO) ModelFromRequest() *model.Exoplanet {
	if req == nil {
		return nil
	}

	return &model.Exoplanet{
		Name:   req.Name,
		StarID: req.StarId,
		Mass:   req.Mass,
		Radius: req.Radius,
		Dist:   req.Dist,
	}
}
