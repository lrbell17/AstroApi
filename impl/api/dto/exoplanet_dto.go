package dto

import (
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/model"
)

type (
	ExoplanetDTO struct {
		ID     uint          `json:"id"`
		Name   string        `json:"name"`
		Mass   MeasuredValue `json:"mass"`
		Radius MeasuredValue `json:"radius"`
		Dist   MeasuredValue `json:"distance"`
		Star   PlanetStarDTO `json:"star"`
	}
	PlanetStarDTO struct {
		ID     uint          `json:"id"`
		Name   string        `json:"name"`
		Mass   MeasuredValue `json:"mass"`
		Radius MeasuredValue `json:"radius"`
		Temp   MeasuredValue `json:"temperature"`
	}
)

func NewExoplanetDTO(planet *model.Exoplanet, datasourceConf *conf.Datasource) *ExoplanetDTO {

	if planet == nil || datasourceConf == nil {
		return nil
	}

	return &ExoplanetDTO{
		ID:     planet.ID,
		Name:   planet.Name,
		Mass:   asMeasuredValue(planet.Mass, datasourceConf.ExoplanetData.Mass.Unit),
		Radius: asMeasuredValue(planet.Radius, datasourceConf.ExoplanetData.Radius.Unit),
		Dist:   asMeasuredValue(planet.Dist, datasourceConf.ExoplanetData.Dist.Unit),
		Star: PlanetStarDTO{
			ID:     planet.Star.ID,
			Name:   planet.Star.Name,
			Mass:   asMeasuredValue(planet.Star.Mass, datasourceConf.StarData.Mass.Unit),
			Radius: asMeasuredValue(planet.Star.Radius, datasourceConf.StarData.Radius.Unit),
			Temp:   asMeasuredValue(planet.Star.Temp, datasourceConf.StarData.Temp.Unit),
		},
	}
}
