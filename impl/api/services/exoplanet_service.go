package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto/request"
	"github.com/lrbell17/astroapi/impl/api/dto/response"
	"github.com/lrbell17/astroapi/impl/cache"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/persistence/repos"
	log "github.com/sirupsen/logrus"
)

type ExoplanetService struct {
	repo     *repos.ExoplanetRepo
	starRepo *repos.StarRepo
	config   *conf.Config
}

// Constructor for exoplanet service
func NewExoplanetService(repo *repos.ExoplanetRepo, starRepo *repos.StarRepo) *ExoplanetService {
	config, err := conf.GetConfig()
	if err != nil {
		log.Fatalf("Unable to start Exoplanet Service: %v", err)
	}

	return &ExoplanetService{repo, starRepo, config}
}

// Call on repo to get the exoplanet by ID and return an ExoplanetDTO
func (s *ExoplanetService) GetById(id uint) (*response.ExoplanetResponseDTO, error) {

	planetResp := &response.ExoplanetResponseDTO{}

	// Check cache
	cacheKey := planetResp.GetCacheKey(id)
	if err := planetResp.GetCached(cacheKey); err == nil {
		return planetResp, nil
	}

	// Get from DB
	planet, err := s.repo.GetById(id)
	if err != nil {
		log.Warnf("Unable to locate planet with ID %v", id)
		return nil, err
	}

	// Build response
	planetResp.ResponseFromDao(planet, &s.config.Datasource)

	// Store in cache
	cache.PutCache(planetResp, cacheKey)

	return planetResp, nil
}

// Call repo to add planet to DB
func (s *ExoplanetService) AddPlanet(planetReq *request.ExoplanetRequestDTO) (*response.ExoplanetResponseDTO, error) {
	planet := planetReq.DaoFromRequest()

	// Check if star is present in DB
	star, err := s.starRepo.GetById(planet.StarID)
	if err != nil {
		log.Warnf("Unable to locate star for planet %v: %v", planet.Name, err)
		return nil, err
	}

	// Insert to DB
	planet, err = s.repo.Insert(planet)
	if err != nil {
		log.Errorf("Error adding planet %v to database: %v", planet.Name, err)
		return nil, err
	}

	// Build response
	planet.Star = *star
	planetResp := &response.ExoplanetResponseDTO{}
	planetResp.ResponseFromDao(planet, &s.config.Datasource)
	star.AddExoplanet(planet)
	starResp := &response.StarResponseDTO{}
	starResp.ResponseFromDao(star, &s.config.Datasource)

	// Add to cache
	cache.PutCache(planetResp, planetResp.GetCacheKey(planetResp.ID))
	cache.PutCache(starResp, starResp.GetCacheKey(starResp.ID))

	return planetResp, nil
}

// Get all habitable planets
func (s *ExoplanetService) GetHabitable() (uint8, []response.ExoplanetResponseDTO, error) {

	var habitablePlanets []response.ExoplanetResponseDTO

	stars, err := s.starRepo.GetAll()
	if err != nil {
		log.Errorf("Error getting stars from database: %v", err)
		return 0, nil, err
	}

	var habitableCount uint8
	for _, star := range stars {
		for _, exoplanet := range star.Exoplanets {
			if exoplanet.IsInHabitableZone(&star) {
				planetResp := &response.ExoplanetResponseDTO{}
				planetResp.ResponseFromDao(&exoplanet, &s.config.Datasource)
				habitablePlanets = append(habitablePlanets, *planetResp)
				habitableCount++
			}
		}
	}
	return habitableCount, habitablePlanets, nil
}
