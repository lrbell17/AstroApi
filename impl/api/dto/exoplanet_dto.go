package dto

import "github.com/lrbell17/astroapi/impl/model"

type ExoplanetDTO struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Mass   float32 `json:"mass"`
	Radius float32 `json:"radius"`
	Dist   float32 `json:"distance"`
	Star   StarDTO `json:"star"`
}

func NewExoplanetDTO(planet *model.Exoplanet) *ExoplanetDTO {
	return &ExoplanetDTO{
		ID:     planet.ID,
		Name:   planet.Name,
		Mass:   planet.Mass,
		Radius: planet.Radius,
		Dist:   planet.Dist,
		Star: StarDTO{
			ID:     planet.Star.ID,
			Name:   planet.Star.Name,
			Mass:   planet.Star.Mass,
			Radius: planet.Star.Radius,
			Temp:   planet.Star.Temp,
		},
	}
}
