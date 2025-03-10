package model

import (
	"fmt"

	"github.com/lrbell17/astroapi/impl/conf"
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

func (*Star) GetTableName() string {
	return starTableName
}

func (*Star) ValidateColumns(header map[string]int) error {
	// TODO
	return nil
}

func (*Star) ParseModel(record []string, colIndices map[string]int, config conf.Datasource) AstroModel {
	// TODO
	return nil
}

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
