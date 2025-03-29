package database

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/persistence/dao"
	log "github.com/sirupsen/logrus"
)

func InitDb() (err error) {

	// Create tables
	err = createTable(&dao.Exoplanet{})
	if err != nil {
		return
	}
	err = createTable(&dao.Star{})
	if err != nil {
		return
	}

	// Import from CSV
	err = importFromCSV(&dao.Star{})
	if err != nil {
		return
	}
	err = importFromCSV(&dao.Exoplanet{})
	if err != nil {
		return
	}

	return
}

// Create SQL table from model
func createTable(dao dao.AstroDAO) (err error) {

	tableName := dao.GetTableName()
	log.Infof("Creating table %v", tableName)

	err = DB.AutoMigrate(dao)
	if err != nil {
		log.Errorf("Unable to create table %v: %v", tableName, err)
	}

	log.Infof("Table %v created successfully", tableName)
	return
}

func importFromCSV(astroDao dao.AstroDAO) error {

	config, _ := conf.GetConfig()
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
	err = astroDao.ValidateColumns(colIndices)
	if err != nil {
		return err
	}

	// Load data in batches
	log.Infof("Loading data for %v", astroDao.GetTableName())
	batchSize := config.Database.Performance.Batchsize
	lineNum, errorCount, successCount := 1, 0, 0
	batch := make([]dao.AstroDAO, 0, batchSize)
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

		recordObj := astroDao.ParseFromCSV(record, colIndices, config.Datasource)
		batch = append(batch, recordObj)

		// Insert batch
		if len(batch) >= batchSize {
			rowsInserted, err := astroDao.CreateBatch(DB, batch)
			if err != nil {
				log.Warnf("Error inserting batch: %v", err)
				errorCount += len(batch) - rowsInserted
				successCount += rowsInserted
			} else {
				successCount += rowsInserted
			}
			// reset batch
			batch = make([]dao.AstroDAO, 0, batchSize)
		}

	}

	// Insert remaining rows
	if len(batch) > 0 {
		rowsInserted, err := astroDao.CreateBatch(DB, batch)
		if err != nil {
			log.Warnf("Error inserting batch: %v", err)
			errorCount += len(batch) - rowsInserted
			successCount += rowsInserted
		} else {
			successCount += rowsInserted
		}
	}

	log.Infof("Import completed: %d records inserted successfully, %d errors", successCount, errorCount)
	return nil
}
