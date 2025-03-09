package model

type AstroModel interface {
	GetModelName() string
	ValidateColumns(map[string]int) error
}
