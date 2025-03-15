package dto

type MeasuredValue struct {
	Value float32 `json:"value"`
	Unit  string  `json:"unit"`
}

func AsMeasuredValue(value float32, unit string) MeasuredValue {
	return MeasuredValue{value, unit}
}
