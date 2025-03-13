package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func Connect() {
	config, _ := conf.GetConfig()
	addr := fmt.Sprintf("%v:%v", config.Cache.Host, config.Cache.Port)
	maxRetries := config.Cache.Performance.MaxRetries
	retryInterval := config.Cache.Performance.RetryInterval

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Cache.Password,
		DB:       0,
	})

	ctx := context.Background()

	// Test connection with retries
	var err error
	for i := 0; i < maxRetries; i++ {
		_, err = RedisClient.Ping(ctx).Result()
		if err == nil {
			log.Info("Redis connection successful")
			return
		}

		time.Sleep(time.Duration(retryInterval) * time.Second)
	}
	log.Fatalf("Failed to connect to Redis: %v", err)
}
