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
	// DTO for star data
	StarDTO struct {
		ID      uint            `json:"id"`
		Name    string          `json:"name"`
		Mass    MeasuredValue   `json:"mass"`
		Radius  MeasuredValue   `json:"radius"`
		Temp    MeasuredValue   `json:"temperature"`
		Planets []StarPlanetDTO `json:"planets"`
	}
	// DTO for exoplanet data nested within star data
	StarPlanetDTO struct {
		ID     uint          `json:"id"`
		Name   string        `json:"name"`
		Mass   MeasuredValue `json:"mass"`
		Radius MeasuredValue `json:"radius"`
		Dist   MeasuredValue `json:"distance"`
	}
)

// Constructor for Star DTO
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

// Get cache key for Star DTO by ID
func (e *StarDTO) GetCacheKey(id uint) string {
	return fmt.Sprintf("star:%d", id)
}

// Get Star DTO from cache by key
func (s *StarDTO) GetCached(cacheKey string) error {
	if s == nil {
		err := fmt.Errorf("star DTO is nil")
		log.Errorf("Could not get cached value for %v: %v", cacheKey, err)
		return err
	}

	ctx := context.Background()

	cached, err := cache.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var starDTO StarDTO
		if err := json.Unmarshal([]byte(cached), &starDTO); err == nil {
			log.Infof("Cache hit on %v", cacheKey)
			*s = starDTO
			return nil
		} else {
			log.Warnf("Unable to unmarshal result for key %v: %v", cacheKey, err)
		}
	}
	log.Infof("Cache miss on %v", cacheKey)
	return fmt.Errorf("cache miss")
}
