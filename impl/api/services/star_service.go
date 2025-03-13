package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto"
	"github.com/lrbell17/astroapi/impl/api/repos"
	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
)

type StarService struct {
	repo   *repos.StarRepo
	config *conf.Config
}

// Constructor for star service
func NewStarService(repo *repos.StarRepo) *StarService {
	config, err := conf.GetConfig()
	if err != nil {
		log.Fatalf("Unable to start Exoplanet Service: %v", err)
	}
	return &StarService{repo, config}
}

// Call on repo to get the star by ID and return an StarDTO
func (s *StarService) GetById(id uint) (*dto.StarDTO, error) {
	star, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewStarDTO(star, &s.config.Datasource), nil
}
