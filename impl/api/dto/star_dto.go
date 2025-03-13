package dto

import "github.com/lrbell17/astroapi/impl/model"

type (
	StarDTO struct {
		ID      uint            `json:"id"`
		Name    string          `json:"name"`
		Mass    float32         `json:"mass"`
		Radius  float32         `json:"radius"`
		Temp    float32         `json:"temperature"`
		Planets []StarPlanetDTO `json:"planets"`
	}
	StarPlanetDTO struct {
		ID     uint    `json:"id"`
		Name   string  `json:"name"`
		Mass   float32 `json:"mass"`
		Radius float32 `json:"radius"`
		Dist   float32 `json:"distance"`
	}
)

func NewStarDTO(star *model.Star) *StarDTO {
	planets := make([]StarPlanetDTO, len(star.Exoplanets))
	for i, planet := range star.Exoplanets {
		planets[i] = StarPlanetDTO{
			ID:     planet.ID,
			Name:   planet.Name,
			Mass:   planet.Mass,
			Radius: planet.Radius,
			Dist:   planet.Dist,
		}
	}

	return &StarDTO{
		ID:      star.ID,
		Name:    star.Name,
		Mass:    star.Mass,
		Radius:  star.Radius,
		Temp:    star.Temp,
		Planets: planets,
	}

}
