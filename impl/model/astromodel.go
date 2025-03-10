package model

import (
	"strconv"

	"github.com/lrbell17/astroapi/impl/conf"
	"gorm.io/gorm"
)

type AstroModel interface {
	GetTableName() string
	ValidateColumns(map[string]int) error
	ParseModel(record []string, colIndices map[string]int, config conf.Datasource) AstroModel
	CreateBatch(db *gorm.DB, batch []AstroModel) error
}

// Gets the string value of a CSV column
func GetStringValue(record []string, colIndices map[string]int, colName string) string {
	if idx, exists := colIndices[colName]; exists && idx < len(record) {
		return record[idx]
	}
	return ""
}

// Gets the float value of a CSV column
func GetFloatValue(record []string, colIndices map[string]int, colName string) float32 {
	if idx, exists := colIndices[colName]; exists && idx < len(record) {
		return ParseFloat(record[idx])
	}
	return 0.0
}

func ParseFloat(val string) float32 {
	if val == "" {
		return 0.0
	}

	floatVal, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return 0.0
	}

	return float32(floatVal)
}
