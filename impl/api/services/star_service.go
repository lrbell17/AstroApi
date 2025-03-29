package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto/request"
	"github.com/lrbell17/astroapi/impl/api/dto/response"
	"github.com/lrbell17/astroapi/impl/cache"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/persistence/repos"
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

	starResp := &response.StarResponseDTO{}

	// Check cache
	cacheKey := starResp.GetCacheKey(id)
	if err := starResp.GetCached(cacheKey); err == nil {
		return starResp, nil
	}

	// Get from DB
	star, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	starResp.ResponseFromDao(star, &s.config.Datasource)

	// Add to cache
	cache.PutCache(starResp, cacheKey)

	return starResp, nil
}

// Call repo to add star to DB
func (s *StarService) AddStar(starReq *request.StarRequestDTO) (*response.StarResponseDTO, error) {

	star := starReq.DaoFromRequest()

	// Insert to DB
	star, err := s.repo.Insert(star)
	if err != nil {
		return nil, err
	}
	// Add to cache
	starResp := &response.StarResponseDTO{}
	starResp.ResponseFromDao(star, &s.config.Datasource)

	cache.PutCache(starResp, starResp.GetCacheKey(starResp.ID))

	return starResp, nil

}
