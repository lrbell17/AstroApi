package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto"
	"github.com/lrbell17/astroapi/impl/api/repos"
	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
)

type ExoplanetService struct {
	repo   *repos.ExoplanetRepo
	config *conf.Config
}

// Constructor for exoplanet service
func NewExoplanetService(repo *repos.ExoplanetRepo) *ExoplanetService {
	config, err := conf.GetConfig()
	if err != nil {
		log.Fatalf("Unable to start Exoplanet Service: %v", err)
	}

	return &ExoplanetService{repo, config}
}

// Call on repo to get the exoplanet by ID and return an ExoplanetDTO
func (s *ExoplanetService) GetById(id uint) (*dto.ExoplanetDTO, error) {
	planet, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewExoplanetDTO(planet, &s.config.Datasource), nil
}
