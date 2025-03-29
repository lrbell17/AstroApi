package request

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin/binding"
)

// Generic request interface
type Request[T any] interface {
	DaoFromRequest() *T
}

// Construct Request DTO from JSON body
func ApplyJsonValues[T any](req Request[T], body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return err
	}
	if err := binding.Validator.ValidateStruct(req); err != nil {
		return err
	}

	return nil
}
