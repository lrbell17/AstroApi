package request

import (
	"encoding/json"
	"io"

	"github.com/lrbell17/astroapi/impl/model"
)

type (
	// DTO for star request
	StarRequestDTO struct {
		Name   string  `json:"name" binding:"required"`
		Mass   float32 `json:"mass"`
		Radius float32 `json:"radius"`
		Temp   float32 `json:"temp"`
	}
)

// Get Request DTO from json
func (req *StarRequestDTO) ApplyJsonValues(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&req)
}

// Get Star from StarRequestDTO
func (req *StarRequestDTO) StarFromRequest() *model.Star {
	if req == nil {
		return nil
	}

	return &model.Star{
		Name:   req.Name,
		Mass:   req.Mass,
		Radius: req.Radius,
		Temp:   req.Temp,
	}
}
