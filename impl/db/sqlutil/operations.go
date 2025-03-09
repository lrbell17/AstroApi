package sqlutil

import (
	"fmt"
	"time"

	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const maxRetries = 5
const retryInterval = 1

// Get a connection to the DB
func GetConnection() (db *gorm.DB, err error) {

	conf, _ := conf.GetConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Name,
	)

	// Try to connect with retries
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

		if err == nil {
			return
		}
		time.Sleep(retryInterval * time.Second)
	}

	log.Errorf("Failed to connect to database: %v", err)
	return
}

// Insert records in batches
func InsertBatch(db *gorm.DB, model []model.AstroModel) error {
	result := db.CreateInBatches(model, len(model))
	return result.Error
}
