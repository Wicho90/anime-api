package season

import (
	"github.com/wicho90/anime-api/internal/entities"
	"strings"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() ([]*entities.Season, error) {
	seasons, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return seasons, nil
}

func (s *service) GetById(id uint64) (*entities.Season, error) {
	season, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return season, err
}

func (s *service) Create(season *entities.Season) (*entities.Season, error) {

	season.Name = strings.TrimSpace(strings.ToLower(season.Name))
	season.Slug = strings.ReplaceAll(season.Name, " ", "-")

	err := s.repository.Create(season)
	if err != nil {
		return nil, err
	}
	return season, err
}

func (s *service) Update(id uint64, season entities.Season) (entities.Season, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}
