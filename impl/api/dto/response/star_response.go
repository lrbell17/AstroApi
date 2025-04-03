package response

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lrbell17/astroapi/impl/api/dto"
	"github.com/lrbell17/astroapi/impl/cache"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/persistence/dao"
	log "github.com/sirupsen/logrus"
)

type (
	// Star response interface
	StarResponse interface {
		Response[dao.Star]
	}
	// DTO for star response
	StarResponseDTO struct {
		ID                 uint              `json:"id"`
		Name               string            `json:"name" binding:"required"`
		Mass               dto.MeasuredValue `json:"mass"`
		Radius             dto.MeasuredValue `json:"radius"`
		Temp               dto.MeasuredValue `json:"temp"`
		Luminosity         dto.MeasuredValue `json:"luminosity"`
		HabitableZoneLower dto.MeasuredValue `json:"habitable_zone_lower_bound"`
		HabitableZoneUpper dto.MeasuredValue `json:"habitable_zone_upper_bound"`
		Planets            []StarPlanetDTO   `json:"planets"`
	}
	// DTO for exoplanet data nested within star response
	StarPlanetDTO struct {
		ID     uint              `json:"id"`
		Name   string            `json:"name"`
		Mass   dto.MeasuredValue `json:"mass"`
		Radius dto.MeasuredValue `json:"radius"`
		Dist   dto.MeasuredValue `json:"distance"`
	}
)

// Get star response DTO from star DAO
func (resp *StarResponseDTO) ResponseFromDao(star *dao.Star, datasourceConf *conf.Datasource) {

	if resp == nil || star == nil || datasourceConf == nil {
		return
	}

	planets := make([]StarPlanetDTO, len(star.Exoplanets))
	for i, planet := range star.Exoplanets {
		planets[i] = StarPlanetDTO{
			ID:     planet.ID,
			Name:   planet.Name,
			Mass:   dto.AsMeasuredValue(planet.Mass, datasourceConf.ExoplanetData.Mass.Unit),
			Radius: dto.AsMeasuredValue(planet.Radius, datasourceConf.ExoplanetData.Radius.Unit),
			Dist:   dto.AsMeasuredValue(planet.Dist, datasourceConf.ExoplanetData.Dist.Unit),
		}
	}

	resp.ID = star.ID
	resp.Name = star.Name
	resp.Mass = dto.AsMeasuredValue(star.Mass, datasourceConf.StarData.Mass.Unit)
	resp.Radius = dto.AsMeasuredValue(star.Radius, datasourceConf.StarData.Radius.Unit)
	resp.Temp = dto.AsMeasuredValue(star.Temp, datasourceConf.StarData.Temp.Unit)
	resp.Luminosity = dto.AsMeasuredValue(star.Luminosity, dao.LuminosityUnits)
	resp.HabitableZoneLower = dto.AsMeasuredValue(star.HabitableZoneLower, dao.HabitableZoneUnits)
	resp.HabitableZoneUpper = dto.AsMeasuredValue(star.HabitableZoneUpper, dao.HabitableZoneUnits)
	resp.Planets = planets

}

// Get cache key for Star DTO by ID
func (e *StarResponseDTO) GetCacheKey(id any) string {
	return fmt.Sprintf("star:%d", id)
}

// Get Star DTO from cache by key
func (s *StarResponseDTO) GetCached(cacheKey string) error {
	if s == nil {
		err := fmt.Errorf("star DTO is nil")
		log.Errorf("Could not get cached value for %v: %v", cacheKey, err)
		return err
	}

	ctx := context.Background()

	cached, err := cache.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var starDTO StarResponseDTO
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
