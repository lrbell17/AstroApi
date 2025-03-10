package model

import (
	"fmt"

	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const planetTableName = "exoplanets"

type (
	Exoplanet struct {
		ID     uint `gorm:"primaryKey"`
		Name   string
		Host   string
		Mass   float32
		Radius float32
		Dist   float32
	}
)

// Get exoplanet table name
func (*Exoplanet) GetTableName() string {
	return planetTableName
}

// Validate CSV columns for exoplanets
func (*Exoplanet) ValidateColumns(header map[string]int) error {

	conf, _ := conf.GetConfig()
	requiredCols := []string{
		conf.Datasource.ExoplanetData.NameCol,
		conf.Datasource.ExoplanetData.HostCol,
		conf.Datasource.ExoplanetData.MassCol,
		conf.Datasource.ExoplanetData.RadiusCol,
		conf.Datasource.ExoplanetData.DistCol,
	}
	for _, col := range requiredCols {
		if _, ok := header[col]; !ok {
			return fmt.Errorf("required column %v does not exist", col)
		}
	}
	log.Infof("Column validation successful: %T %v", Exoplanet{}, requiredCols)
	return nil
}

// Parse exoplanet from CSV records
func (*Exoplanet) ParseModel(record []string, colIndices map[string]int, config conf.Datasource) AstroModel {

	explanetDataConf := config.ExoplanetData
	return &Exoplanet{
		Name:   GetStringValue(record, colIndices, explanetDataConf.NameCol),
		Host:   GetStringValue(record, colIndices, explanetDataConf.HostCol),
		Dist:   GetFloatValue(record, colIndices, explanetDataConf.DistCol),
		Radius: GetFloatValue(record, colIndices, explanetDataConf.RadiusCol),
		Mass:   GetFloatValue(record, colIndices, explanetDataConf.MassCol),
	}
}

// Insert batch of exoplanets into DB
func (e *Exoplanet) CreateBatch(db *gorm.DB, batch []AstroModel) error {

	batchSize := len(batch)
	if batchSize == 0 {
		return fmt.Errorf("empty batch")
	}

	exoplanets := make([]*Exoplanet, 0, len(batch))
	for _, record := range batch {
		if exo, ok := record.(*Exoplanet); ok {
			exoplanets = append(exoplanets, exo)
		}
	}
	return db.Model(&Exoplanet{}).CreateInBatches(exoplanets, batchSize).Error
}
