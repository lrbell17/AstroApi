package db

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/db/sqlutil"
	"github.com/lrbell17/astroapi/impl/model"
	exoplanet "github.com/lrbell17/astroapi/impl/model/exoplanet"
	stars "github.com/lrbell17/astroapi/impl/model/star"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() (err error) {

	db, err = sqlutil.GetConnection()
	if err != nil {
		return fmt.Errorf("unable to get DB connection: %v", err)
	}

	// Create tables
	err = createTable(&exoplanet.Exoplanet{})
	if err != nil {
		return
	}
	err = createTable(&stars.Star{})
	if err != nil {
		return
	}

	// Import from CSV
	err = importFromCSV(&exoplanet.Exoplanet{})
	if err != nil {
		return
	}

	return
}

// Create SQL table from model
func createTable(model model.AstroModel) (err error) {

	modelName := model.GetModelName()
	log.Infof("Creating table for %v model", modelName)

	err = db.AutoMigrate(model)
	if err != nil {
		log.Errorf("Unable to create table for %v model: %v", modelName, err)
	}

	log.Infof("Table created successfully for %v model", modelName)
	return
}

func importFromCSV(model model.AstroModel) error {
	conf, _ := conf.GetConfig()
	filePath := conf.Datasource.File

	// open CSV
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v: %v", filePath, err)
	}
	defer file.Close()

	// Reader for CSV
	reader := csv.NewReader(file)
	reader.Comment = '#'

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header for %v: %v", filePath, err)
	}

	// Map column name to indexes
	colIndices := make(map[string]int)
	for idx, col := range header {
		colIndices[col] = idx
	}

	// Validate required columns
	err = model.ValidateColumns(colIndices)
	if err != nil {
		return err
	}

	return nil
}
