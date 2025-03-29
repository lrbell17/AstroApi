package dao

import "strconv"

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

// Parse float from strings
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
