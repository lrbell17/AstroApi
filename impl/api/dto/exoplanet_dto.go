package dto

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lrbell17/astroapi/impl/cache"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/model"
	log "github.com/sirupsen/logrus"
)

type (
	// DTO for exoplanet data
	ExoplanetDTO struct {
		ID     uint          `json:"id"`
		Name   string        `json:"name"`
		Mass   MeasuredValue `json:"mass"`
		Radius MeasuredValue `json:"radius"`
		Dist   MeasuredValue `json:"distance"`
		Star   PlanetStarDTO `json:"star"`
	}
	// DTO for star data nested within exoplanet data
	PlanetStarDTO struct {
		ID     uint          `json:"id"`
		Name   string        `json:"name"`
		Mass   MeasuredValue `json:"mass"`
		Radius MeasuredValue `json:"radius"`
		Temp   MeasuredValue `json:"temperature"`
	}
)

// Constructor for Exoplanet DTO
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

// Get cache key for Exoplanet DTO by ID
func (e *ExoplanetDTO) GetCacheKey(id uint) string {
	return fmt.Sprintf("exoplanet:%d", id)
}

// Get Exoplanet DTO from cache by key
func (e *ExoplanetDTO) GetCached(cacheKey string) error {
	if e == nil {
		err := fmt.Errorf("exoplanet DTO is nil")
		log.Errorf("Could not get cached value for %v: %v", cacheKey, err)
		return err
	}

	ctx := context.Background()

	cached, err := cache.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var planetDTO ExoplanetDTO
		if err := json.Unmarshal([]byte(cached), &planetDTO); err == nil {
			log.Infof("Cache hit on %v", cacheKey)
			*e = planetDTO
			return nil
		} else {
			log.Warnf("Unable to unmarshal result for key %v: %v", cacheKey, err)
		}
	}
	log.Infof("Cache miss on %v", cacheKey)
	return fmt.Errorf("cache miss")
}
