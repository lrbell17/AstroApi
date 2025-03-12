package dto

type StarDTO struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Mass   float32 `json:"mass"`
	Radius float32 `json:"radius"`
	Temp   float32 `json:"temperature"`
}
