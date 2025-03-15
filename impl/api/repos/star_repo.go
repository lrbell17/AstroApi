package repos

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lrbell17/astroapi/impl/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	ErrStarExists = "star already exists"
)

type StarRepo struct {
	db *gorm.DB
}

// Constructor for star repo
func NewStarRepo(db *gorm.DB) *StarRepo {
	return &StarRepo{db}
}

// Get star by ID from DB
func (r *StarRepo) GetById(id uint) (*model.Star, error) {
	var star model.Star
	result := r.db.Preload("Exoplanets").First(&star, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &star, nil
}

// Add star to DB
func (r *StarRepo) Insert(s *model.Star) (*model.Star, error) {

	if err := r.db.Create(s).Error; err != nil {

		// duplicate key error
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			log.Warnf("%s: %v", ErrStarExists, err)
			return nil, errors.New(ErrStarExists)
		}

		log.Errorf("Error adding star: %v", err)
		return nil, err
	}

	return s, nil
}
