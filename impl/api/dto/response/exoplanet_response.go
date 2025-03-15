package response

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lrbell17/astroapi/impl/api/dto"
	"github.com/lrbell17/astroapi/impl/cache"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/model"
	log "github.com/sirupsen/logrus"
)

type (
	// DTO for exoplanet data
	ExoplanetDTO struct {
		ID     uint              `json:"id"`
		Name   string            `json:"name"`
		Mass   dto.MeasuredValue `json:"mass"`
		Radius dto.MeasuredValue `json:"radius"`
		Dist   dto.MeasuredValue `json:"distance"`
		Star   PlanetStarDTO     `json:"star"`
	}
	// DTO for star data nested within exoplanet data
	PlanetStarDTO struct {
		ID     uint              `json:"id"`
		Name   string            `json:"name"`
		Mass   dto.MeasuredValue `json:"mass"`
		Radius dto.MeasuredValue `json:"radius"`
		Temp   dto.MeasuredValue `json:"temp"`
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
		Mass:   dto.AsMeasuredValue(planet.Mass, datasourceConf.ExoplanetData.Mass.Unit),
		Radius: dto.AsMeasuredValue(planet.Radius, datasourceConf.ExoplanetData.Radius.Unit),
		Dist:   dto.AsMeasuredValue(planet.Dist, datasourceConf.ExoplanetData.Dist.Unit),
		Star: PlanetStarDTO{
			ID:     planet.Star.ID,
			Name:   planet.Star.Name,
			Mass:   dto.AsMeasuredValue(planet.Star.Mass, datasourceConf.StarData.Mass.Unit),
			Radius: dto.AsMeasuredValue(planet.Star.Radius, datasourceConf.StarData.Radius.Unit),
			Temp:   dto.AsMeasuredValue(planet.Star.Temp, datasourceConf.StarData.Temp.Unit),
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
