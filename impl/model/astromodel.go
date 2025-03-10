package model

import (
	"github.com/lrbell17/astroapi/impl/conf"
	"gorm.io/gorm"
)

type AstroModel interface {
	GetTableName() string
	ValidateColumns(map[string]int) error
	ParseModel(record []string, colIndices map[string]int, config conf.Datasource) AstroModel
	CreateBatch(db *gorm.DB, batch []AstroModel) error
}
