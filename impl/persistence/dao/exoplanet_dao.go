package dao

import (
	"fmt"

	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const planetTableName = "exoplanets"

type (
	Exoplanet struct {
		ID            uint   `gorm:"primaryKey"`
		Name          string `gorm:"uniqueIndex:idx_name_star"`
		Host          string
		StarID        uint `gorm:"index;uniqueIndex:idx_name_star"`
		Star          Star `gorm:"foreignKey:StarID"`
		Mass          float32
		Radius        float32
		Dist          float32
		OrbitalPeriod float32
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
		conf.Datasource.ExoplanetData.OrbitalPeriod.ColName,
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
func (*Exoplanet) ParseFromCSV(record []string, colIndices map[string]int, config conf.Datasource) AstroDAO {

	exoplanetDataConf := config.ExoplanetData
	return &Exoplanet{
		Name:          GetStringValue(record, colIndices, exoplanetDataConf.Name.ColName),
		Host:          GetStringValue(record, colIndices, exoplanetDataConf.Host.ColName),
		Dist:          GetFloatValue(record, colIndices, exoplanetDataConf.Dist.ColName),
		Radius:        GetFloatValue(record, colIndices, exoplanetDataConf.Radius.ColName),
		Mass:          GetFloatValue(record, colIndices, exoplanetDataConf.Mass.ColName),
		OrbitalPeriod: GetFloatValue(record, colIndices, exoplanetDataConf.OrbitalPeriod.ColName),
	}
}

// Insert batch of exoplanets into DB
func (e *Exoplanet) CreateBatch(db *gorm.DB, batch []AstroDAO) (int, error) {

	batchSize := len(batch)
	if batchSize == 0 {
		return 0, fmt.Errorf("empty batch")
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
	result := db.Model(&Exoplanet{}).CreateInBatches(exoplanets, batchSize)

	return int(result.RowsAffected), result.Error
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
