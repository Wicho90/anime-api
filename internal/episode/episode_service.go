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
		return nil, fmt.Errorf("failed to get episode: %w", err)
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
func (s *service) Create(episode *entities.Episode) (*entities.Episode, error) {
	/*name := strings.ToLower(createDto.Name)
	episode := &Episode{
		Name:          name,
		EpisodeNumber: createDto.EpisodeNumber,
		Slug:          strings.ReplaceAll(strings.TrimSpace(name), " ", "-"),
	}
	*/
	seasonFound, err := s.seasonRepository.GetById(episode.SeasonId)
	if err != nil {
		return nil, err
	}

	episode.Name = strings.TrimSpace(strings.ToLower(episode.Name))

	episode.Slug = seasonFound.Slug + "-" + strconv.Itoa(int(episode.Number))

	err = s.repository.Create(episode)
	if err != nil {
		return nil, fmt.Errorf("failed to create episode: %w", err)
	}

	return episode, nil
}

func (s *service) Update(id uint64, episode *entities.Episode) (*entities.Episode, error) {
	// TODO hacer transacción
	_, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if _, err := s.seasonRepository.GetById(episode.SeasonId); err != nil {
		return nil, err
	}

	episode.ID = id
	episode.Name = strings.TrimSpace(strings.ToLower(episode.Name))
	episode.Slug = strings.ReplaceAll(episode.Name, " ", "-")

	err = s.repository.Update(episode)
	if err != nil {
		return nil, err
	}

	return episode, nil
}

func (s *service) Delete(id uint64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
