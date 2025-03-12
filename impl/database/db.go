package database

import (
	"fmt"
	"time"

	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Get a connection to the DB
func Connect() {

	conf, _ := conf.GetConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Name,
	)

	maxRetries := conf.Database.Performance.MaxRetries
	retryInterval := conf.Database.Performance.RetryInterval

	// Try to connect with retries
	var err error
	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

		if err == nil {
			return
		}
		time.Sleep(time.Duration(retryInterval) * time.Second)
	}

	log.Fatalf("Failed to connect to database: %v", err)
}
