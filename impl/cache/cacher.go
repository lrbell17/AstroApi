package cache

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type (
	Cacher interface {
		GetCached(cacheKey string) error
	}
)

// Adds item to cache
func PutCache(c Cacher, cacheKey string) {
	if c == nil {
		log.Errorf("Cannot add nil pointer to cache")
		return
	}

	log.Infof("Adding %v to cache for %v", cacheKey, Expiry)

	ctx := context.Background()

	data, err := json.Marshal(c)
	if err != nil {
		log.Warnf("Unable to marshall %v: %v", cacheKey, err)
	}
	err = RedisClient.Set(ctx, cacheKey, data, Expiry).Err()
	if err != nil {
		log.Warnf("Unable to add %v to cache: %v", cacheKey, err)
	}
}
