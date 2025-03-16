package request

import (
	"github.com/lrbell17/astroapi/impl/model"
)

type (
	// Star request interface
	StarRequest interface {
		Request[model.Star]
	}
	// DTO for star request
	StarRequestDTO struct {
		Name   string  `json:"name" binding:"required"`
		Mass   float32 `json:"mass"`
		Radius float32 `json:"radius"`
		Temp   float32 `json:"temp"`
	}
)

// Get Star model from Star request DTO
func (req *StarRequestDTO) ModelFromRequest() *model.Star {
	if req == nil {
		return nil
	}

	return &model.Star{
		Name:   req.Name,
		Mass:   req.Mass,
		Radius: req.Radius,
		Temp:   req.Temp,
	}
}
