package dao

import (
	"fmt"

	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const starTableName = "stars"

type (
	Star struct {
		ID                 uint   `gorm:"primaryKey"`
		Name               string `gorm:"uniqueIndex"`
		Mass               float32
		Radius             float32
		Temp               float32
		Luminosity         float32
		HabitableZoneLower float32
		HabitableZoneUpper float32
		Exoplanets         []Exoplanet `gorm:"foreignKey:StarID"`
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
		conf.Datasource.StarData.Name.ColName,
		conf.Datasource.StarData.Mass.ColName,
		conf.Datasource.StarData.Radius.ColName,
		conf.Datasource.StarData.Temp.ColName,
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
func (*Star) ParseFromCSV(record []string, colIndices map[string]int, config conf.Datasource) AstroDAO {
	starDataConf := config.StarData
	return &Star{
		Name:   GetStringValue(record, colIndices, starDataConf.Name.ColName),
		Temp:   GetFloatValue(record, colIndices, starDataConf.Temp.ColName),
		Radius: GetFloatValue(record, colIndices, starDataConf.Radius.ColName),
		Mass:   GetFloatValue(record, colIndices, starDataConf.Mass.ColName),
	}
}

// Insert batch of stars into DB
func (e *Star) CreateBatch(db *gorm.DB, batch []AstroDAO) (int, error) {

	batchSize := len(batch)
	if batchSize == 0 {
		return 0, fmt.Errorf("empty batch")
	}

	stars := make([]*Star, 0, len(batch))
	for _, record := range batch {
		if s, ok := record.(*Star); ok {
			s.EnrichFields()
			stars = append(stars, s)
		}
	}

	result := db.Model(&Star{}).
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "name"}}, DoNothing: true}). // Ignore duplicates
		CreateInBatches(stars, batchSize)

	return int(result.RowsAffected), result.Error

}

// Add an exoplanet to a star
func (s *Star) AddExoplanet(e *Exoplanet) {
	if s == nil {
		return
	}

	s.Exoplanets = append(s.Exoplanets, *e)
}
