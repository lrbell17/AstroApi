package cache

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type (
	Cacher interface {
		GetCached(cacheKey string) error
		GetCacheKey(id any) string
	}
)

// Adds item to cache
func PutCache(c Cacher, cacheKey string) {
	go func() {
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
	}()
}

// Invalidate cache for keys matching pattern
func InvalidateCacheKeys(pattern string) {
	go func() {
		log.Infof("Deleting cache keys matching pattern %s", pattern)
		ctx := context.Background()
		iter := RedisClient.Scan(ctx, 0, pattern, 0).Iterator()

		for iter.Next(ctx) {
			if err := RedisClient.Del(ctx, iter.Val()).Err(); err != nil {
				log.Warnf("Failed to delete cache key %s: %v", iter.Val(), err)
			}
		}

		if err := iter.Err(); err != nil {
			log.Errorf("Error scanning cache keys: %v", err)
		}
	}()
}
