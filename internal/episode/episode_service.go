package episode

import (
	"fmt"
	"github.com/wicho90/anime-api/internal/entities"
	"github.com/wicho90/anime-api/internal/season"
	"strconv"
	"strings"
)

type service struct {
	repository       Repository
	seasonRepository season.Repository
}

func NewService(repository Repository, seasonRepository season.Repository) Service {
	return &service{
		repository:       repository,
		seasonRepository: seasonRepository,
	}
}

func (s *service) GetAll() ([]*entities.Episode, error) {
	episodes, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return episodes, nil
}

func (s *service) GetLatest() ([]*entities.EpisodeWithImage, error) {
	episodes, err := s.repository.GetLatest()
	if err != nil {
		return nil, err
	}

	return episodes, nil
}

func (s *service) GetByID(id uint64) (*entities.Episode, error) {
	episode, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return episode, err
}

func (s *service) GetBySlug(slug string) (*entities.EpisodeWithSeasonSlug, error) {
	episode, err := s.repository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}

	return episode, nil
}

func (s *service) Create(episode *entities.Episode) error {
	seasonFound, err := s.seasonRepository.GetById(episode.SeasonId)
	if err != nil {
		return err
	}

	episode.Name = strings.TrimSpace(strings.ToLower(episode.Name))
	episode.Slug = seasonFound.Slug + "-" + strconv.Itoa(int(episode.Number))

	err = s.repository.Create(episode)
	if err != nil {
		return fmt.Errorf("failed to create episode: %w", err)
	}

	return nil
}

func (s *service) Update(id uint64, episode *entities.Episode) error {
	// TODO hacer transacción
	_, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	seasonFound, err := s.seasonRepository.GetById(episode.SeasonId)
	if err != nil {
		return err
	}

	episode.ID = id
	episode.Name = strings.TrimSpace(strings.ToLower(episode.Name))
	episode.Slug = seasonFound.Slug + "-" + strconv.Itoa(int(episode.Number))

	err = s.repository.Update(episode)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(id uint64) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}

	return nil
}
