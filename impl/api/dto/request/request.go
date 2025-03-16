package request

import (
	"encoding/json"
	"io"
)

// Generic request interface
type Request[T any] interface {
	ModelFromRequest() *T
}

// Construct Request DTO from JSON body
func ApplyJsonValues[T any](req Request[T], body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&req)
}
