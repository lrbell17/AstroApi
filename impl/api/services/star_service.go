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
	star.EnrichFields()

	// Insert to DB
	star, err := s.repo.Insert(star)
	if err != nil {
		return nil, err
	}
	// Get response
	starResp := &response.StarResponseDTO{}
	starResp.ResponseFromDao(star, &s.config.Datasource)

	// Add to cache
	cache.PutCache(starResp, starResp.GetCacheKey(starResp.ID))

	// Invalidate cache for star name search
	cache.InvalidateCacheKeys("star_list*")

	return starResp, nil

}

// Call repo to search for star by name
func (s *StarService) SearchByName(name string, limit int) (*response.StarResponseDTOList, error) {
	starResponses := &response.StarResponseDTOList{
		Stars: []response.StarResponseDTO{},
	}

	// Check cache
	cacheKey := starResponses.GetCacheKey(name)
	if err := starResponses.GetCached(cacheKey); err == nil {
		return starResponses, nil
	}

	// Query DB
	stars, err := s.repo.SearchByName(name, limit)
	if err != nil {
		log.Errorf("Unable to complete search for stars by name %v and limit %d: %v", name, limit, err)
		return starResponses, err
	}

	// Get response
	for _, star := range stars {
		starResp := &response.StarResponseDTO{}
		starResp.ResponseFromDao(&star, &s.config.Datasource)
		starResponses.Stars = append(starResponses.Stars, *starResp)
	}

	// Add response to cache
	cache.PutCache(starResponses, cacheKey)

	return starResponses, nil
}
