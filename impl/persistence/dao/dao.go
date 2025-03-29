package dao

import (
	"github.com/lrbell17/astroapi/impl/conf"
	"gorm.io/gorm"
)

type AstroDAO interface {
	GetTableName() string
	ValidateColumns(map[string]int) error
	ParseFromCSV(record []string, colIndices map[string]int, config conf.Datasource) AstroDAO
	CreateBatch(db *gorm.DB, batch []AstroDAO) (int, error)
}
