package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto"
	"github.com/lrbell17/astroapi/impl/api/repos"
	"github.com/lrbell17/astroapi/impl/cache"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/model"
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

	starDTO := &dto.StarDTO{}

	// Check cache
	cacheKey := starDTO.GetCacheKey(id)
	if err := starDTO.GetCached(cacheKey); err == nil {
		return starDTO, nil
	}

	// Get from DB
	star, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	starDTO = dto.NewStarDTO(star, &s.config.Datasource)

	// Add to cache
	cache.PutCache(starDTO, cacheKey)

	return starDTO, nil
}

// Call repo too add star to DB
func (s *StarService) AddStar(star *model.Star) error {

	// Insert to DB
	if err := s.repo.Insert(star); err != nil {
		return err
	}
	// Add to cache
	starDTO := dto.NewStarDTO(star, &s.config.Datasource)
	cache.PutCache(starDTO, starDTO.GetCacheKey(starDTO.ID))

	return nil

}
