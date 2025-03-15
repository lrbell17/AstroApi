package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto/request"
	"github.com/lrbell17/astroapi/impl/api/dto/response"
	"github.com/lrbell17/astroapi/impl/api/repos"
	"github.com/lrbell17/astroapi/impl/cache"
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
func (s *StarService) GetById(id uint) (*response.StarResponseDTO, error) {

	starDTO := &response.StarResponseDTO{}

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
	starDTO = response.ResponseFromStar(star, &s.config.Datasource)

	// Add to cache
	cache.PutCache(starDTO, cacheKey)

	return starDTO, nil
}

// Call repo to add star to DB
func (s *StarService) AddStar(starReq *request.StarRequestDTO) (*response.StarResponseDTO, error) {

	star := starReq.StarFromRequest()

	// Insert to DB
	star, err := s.repo.Insert(star)
	if err != nil {
		return nil, err
	}
	// Add to cache
	starResp := response.ResponseFromStar(star, &s.config.Datasource)
	cache.PutCache(starResp, starResp.GetCacheKey(starResp.ID))

	return starResp, nil

}
