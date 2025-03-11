package db

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	config *conf.Config
)

func InitDb() (err error) {

	// Get DB session
	db, err = GetSession()
	if err != nil {
		return fmt.Errorf("unable to get DB connection: %v", err)
	}

	// Get config
	config, err = conf.GetConfig()
	if err != nil {
		return err
	}

	// Create tables
	err = createTable(&model.Exoplanet{})
	if err != nil {
		return
	}
	err = createTable(&model.Star{})
	if err != nil {
		return
	}

	// Import from CSV
	err = importFromCSV(&model.Star{})
	if err != nil {
		return
	}
	err = importFromCSV(&model.Exoplanet{})
	if err != nil {
		return
	}

	return
}

// Create SQL table from model
func createTable(model model.AstroModel) (err error) {

	tableName := model.GetTableName()
	log.Infof("Creating table %v", tableName)

	err = db.AutoMigrate(model)
	if err != nil {
		log.Errorf("Unable to create table %v: %v", tableName, err)
	}

	log.Infof("Table %v created successfully", tableName)
	return
}

func importFromCSV(astroModel model.AstroModel) error {
	filePath := config.Datasource.File

	// open CSV
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v: %v", filePath, err)
	}
	defer file.Close()

	// Reader header from CSV
	reader := csv.NewReader(file)
	reader.Comment = '#'
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header for %v: %v", filePath, err)
	}

	// Map column name to indexes and validate
	colIndices := make(map[string]int)
	for idx, col := range header {
		colIndices[col] = idx
	}
	err = astroModel.ValidateColumns(colIndices)
	if err != nil {
		return err
	}

	// Load data in batches
	log.Infof("Loading data for %v", astroModel.GetTableName())
	batchSize := config.Database.Performance.Batchsize
	lineNum, errorCount, successCount := 1, 0, 0
	batch := make([]model.AstroModel, 0, batchSize)
	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Warnf("Error reading line %d, skipping: %v", lineNum, err)
			errorCount++
			lineNum++
			continue
		}

		recordObj := astroModel.ParseModel(record, colIndices, config.Datasource)
		batch = append(batch, recordObj)

		// Insert batch
		if len(batch) >= batchSize {
			if err := astroModel.CreateBatch(db, batch); err != nil {
				log.Warnf("Error inserting batch: %v", err)
				errorCount += len(batch)
			} else {
				successCount += len(batch)
			}
			// reset batch
			batch = make([]model.AstroModel, 0, batchSize)
		}

	}

	// Insert remaining rows
	if len(batch) > 0 {
		if err := astroModel.CreateBatch(db, batch); err != nil {
			log.Warnf("Error inserting batch: %v", err)
			errorCount += len(batch)
		} else {
			successCount += len(batch)
		}
	}

	log.Infof("Import completed: %d records inserted successfully, %d errors", successCount, errorCount)
	return nil
}
