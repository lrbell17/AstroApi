package services

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrStarExists   = "star already exists"
	ErrPlanetExists = "planet already exists"
	ErrInvalidId    = "invalid ID"
	ErrInternal     = "internal error"
)

// Checks if an error is a duplicate key error
func IsDuplicateKeyErr(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}
	return false
}
