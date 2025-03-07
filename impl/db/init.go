package db

import (
	"reflect"

	"github.com/lrbell17/astroapi/impl/db/sqlutil"
	exoplanet "github.com/lrbell17/astroapi/impl/model/exoplanet"
	stars "github.com/lrbell17/astroapi/impl/model/star"
	log "github.com/sirupsen/logrus"
)

func InitDb() {
	createTable(&exoplanet.Exoplanet{})
	createTable(&stars.Star{})
}

func createTable(model any) (err error) {

	modelName := reflect.TypeOf(model).Elem().Name()

	log.Infof("Creating table for %v model", modelName)

	db, err := sqlutil.GetConnection()
	if err != nil {
		return
	}

	log.Infof("Creating table for %v model", modelName)
	err = db.AutoMigrate(model)
	if err != nil {
		log.Errorf("Unable to create table for %v model: %v", modelName, err)
	}

	log.Infof("Table created successfully for %v model", modelName)
	return
}
