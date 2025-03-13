package dto

import (
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/model"
)

type (
	StarDTO struct {
		ID      uint            `json:"id"`
		Name    string          `json:"name"`
		Mass    MeasuredValue   `json:"mass"`
		Radius  MeasuredValue   `json:"radius"`
		Temp    MeasuredValue   `json:"temperature"`
		Planets []StarPlanetDTO `json:"planets"`
	}
	StarPlanetDTO struct {
		ID     uint          `json:"id"`
		Name   string        `json:"name"`
		Mass   MeasuredValue `json:"mass"`
		Radius MeasuredValue `json:"radius"`
		Dist   MeasuredValue `json:"distance"`
	}
)

func NewStarDTO(star *model.Star, datasourceConf *conf.Datasource) *StarDTO {

	if star == nil || datasourceConf == nil {
		return nil
	}

	planets := make([]StarPlanetDTO, len(star.Exoplanets))
	for i, planet := range star.Exoplanets {
		planets[i] = StarPlanetDTO{
			ID:     planet.ID,
			Name:   planet.Name,
			Mass:   asMeasuredValue(planet.Mass, datasourceConf.ExoplanetData.Mass.Unit),
			Radius: asMeasuredValue(planet.Radius, datasourceConf.ExoplanetData.Radius.Unit),
			Dist:   asMeasuredValue(planet.Dist, datasourceConf.ExoplanetData.Dist.Unit),
		}
	}

	return &StarDTO{
		ID:      star.ID,
		Name:    star.Name,
		Mass:    asMeasuredValue(star.Mass, datasourceConf.StarData.Mass.Unit),
		Radius:  asMeasuredValue(star.Radius, datasourceConf.StarData.Radius.Unit),
		Temp:    asMeasuredValue(star.Temp, datasourceConf.StarData.Temp.Unit),
		Planets: planets,
	}

}
