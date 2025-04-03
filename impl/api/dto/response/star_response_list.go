package response

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lrbell17/astroapi/impl/cache"
	log "github.com/sirupsen/logrus"
)

type (
	// List of star response DTOs
	StarResponseDTOList struct {
		Stars []StarResponseDTO `json:"stars"`
	}
)

// Get cache key of star response containers
func (s *StarResponseDTOList) GetCacheKey(name any) string {
	return fmt.Sprintf("star_list:%v", name)
}

func (s *StarResponseDTOList) GetCached(cacheKey string) error {
	if s == nil {
		err := fmt.Errorf("star DTO container is nil")
		log.Errorf("Could not get cached value for %v: %v", cacheKey, err)
		return err
	}

	ctx := context.Background()

	cached, err := cache.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var starDTOList StarResponseDTOList
		if err := json.Unmarshal([]byte(cached), &starDTOList); err == nil {
			log.Infof("Cache hit on %v", cacheKey)
			*s = starDTOList
			return nil
		} else {
			log.Warnf("Unable to unmarshal result for key %v: %v", cacheKey, err)
		}
	}
	log.Infof("Cache miss on %v", cacheKey)
	return fmt.Errorf("cache miss")

}
