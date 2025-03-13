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
		StarID uint `gorm:"index"`
		Star   Star `gorm:"foreignKey:StarID"`
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
		conf.Datasource.ExoplanetData.Name.ColName,
		conf.Datasource.ExoplanetData.Host.ColName,
		conf.Datasource.ExoplanetData.Mass.ColName,
		conf.Datasource.ExoplanetData.Radius.ColName,
		conf.Datasource.ExoplanetData.Dist.ColName,
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
		Name:   GetStringValue(record, colIndices, explanetDataConf.Name.ColName),
		Host:   GetStringValue(record, colIndices, explanetDataConf.Host.ColName),
		Dist:   GetFloatValue(record, colIndices, explanetDataConf.Dist.ColName),
		Radius: GetFloatValue(record, colIndices, explanetDataConf.Radius.ColName),
		Mass:   GetFloatValue(record, colIndices, explanetDataConf.Mass.ColName),
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

			// Get the star ID
			if err := exo.EnrichWithStarId(db); err != nil {
				log.Warnf("Failed to get star for %v, skipping. Error: %v", exo, err)
				continue
			}

			exoplanets = append(exoplanets, exo)
		}
	}
	return db.Model(&Exoplanet{}).CreateInBatches(exoplanets, batchSize).Error
}

// Enrich exoplanet with the star ID by star name
func (e *Exoplanet) EnrichWithStarId(db *gorm.DB) error {
	var star Star
	result := db.Where("name = ?", e.Host).First(&star)
	if result.Error != nil {
		return result.Error
	}

	e.StarID = star.ID
	return nil
}
