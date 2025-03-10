package model

import (
	"fmt"

	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const starTableName = "stars"

type (
	Star struct {
		ID     uint `gorm:"primaryKey"`
		Name   string
		Mass   float32
		Radius float32
		Temp   float32
	}
)

// Get stars table name
func (*Star) GetTableName() string {
	return starTableName
}

// Validate CSV columns for star
func (*Star) ValidateColumns(header map[string]int) error {

	conf, _ := conf.GetConfig()
	requiredCols := []string{
		conf.Datasource.StarData.NameCol,
		conf.Datasource.StarData.MassCol,
		conf.Datasource.StarData.RadiusCol,
		conf.Datasource.StarData.TempCol,
	}
	for _, col := range requiredCols {
		if _, ok := header[col]; !ok {
			return fmt.Errorf("required column %v does not exist", col)
		}
	}
	log.Infof("Column validation successful: %T %v", Star{}, requiredCols)
	return nil
}

// Parse star from CSV record
func (*Star) ParseModel(record []string, colIndices map[string]int, config conf.Datasource) AstroModel {
	starDataConf := config.StarData
	return &Star{
		Name:   GetStringValue(record, colIndices, starDataConf.NameCol),
		Temp:   GetFloatValue(record, colIndices, starDataConf.TempCol),
		Radius: GetFloatValue(record, colIndices, starDataConf.RadiusCol),
		Mass:   GetFloatValue(record, colIndices, starDataConf.MassCol),
	}
}

// Insert batch of stars into DB
func (e *Star) CreateBatch(db *gorm.DB, batch []AstroModel) error {

	batchSize := len(batch)
	if batchSize == 0 {
		return fmt.Errorf("empty batch")
	}

	stars := make([]*Star, 0, len(batch))
	for _, record := range batch {
		if exo, ok := record.(*Star); ok {
			stars = append(stars, exo)
		}
	}
	return db.Model(&Star{}).CreateInBatches(stars, batchSize).Error
}
