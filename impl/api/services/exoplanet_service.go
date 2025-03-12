package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto"
	"github.com/lrbell17/astroapi/impl/api/repos"
)

type ExoplanetService struct {
	repo *repos.ExoplanetRepo
}

func NewExoplanetService(repo *repos.ExoplanetRepo) *ExoplanetService {
	return &ExoplanetService{repo}
}

func (s *ExoplanetService) GetByID(id uint) (*dto.ExoplanetDTO, error) {
	planet, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewExoplanetDTO(planet), nil
}
