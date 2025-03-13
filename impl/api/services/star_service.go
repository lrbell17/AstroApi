package services

import (
	"github.com/lrbell17/astroapi/impl/api/dto"
	"github.com/lrbell17/astroapi/impl/api/repos"
)

type StarService struct {
	repo *repos.StarRepo
}

func NewStarService(repo *repos.StarRepo) *StarService {
	return &StarService{repo}
}

func (s *StarService) GetById(id uint) (*dto.StarDTO, error) {
	star, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return dto.NewStarDTO(star), nil
}
