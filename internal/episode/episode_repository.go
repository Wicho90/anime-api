package episode

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/wicho90/anime-api/internal/entities"
	"github.com/wicho90/anime-api/internal/ex"
	"log"
)

const (
	queryGetAll     = "SELECT id, name, number, duration, url, slug, season_id FROM episodes"
	queryGetLatest  = "SELECT e.id, e.name, e.slug, s.image_url FROM episodes as e INNER JOIN seasons as s ON  season_id = s.id"
	queryGetById    = "SELECT id, name, number, duration, url, slug, season_id FROM episodes WHERE id = $1"
	queryCreate     = "INSERT INTO episodes (name, number, duration, url, slug, season_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	queryUpdate     = "UPDATE episodes SET name = $1, number = $2, duration = $3, url = $4, slug = $5, season_id = $6 WHERE id = $7"
	queryDeleteById = "DELETE FROM episodes WHERE id = $1"
	queryGetBySlug  = "SELECT e.id, e.name, e.number, e.duration, e.url, e.slug, s.slug FROM episodes as e INNER JOIN seasons as s ON  season_id = s.id WHERE e.slug= $1"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]*entities.Episode, error) {
	rows, err := r.db.Query(queryGetAll)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %s", err)
		}
	}(rows)

	var episodes []*entities.Episode
	for rows.Next() {
		episode := &entities.Episode{}
		if err := rows.Scan(&episode.ID, &episode.Name, &episode.Number,
			&episode.Duration, &episode.Url, &episode.Slug, &episode.SeasonId); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		episodes = append(episodes, episode)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return episodes, nil
}

func (r *repository) GetLatest() ([]*entities.EpisodeWithImage, error) {
	rows, err := r.db.Query(queryGetLatest)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows:: %s", err)
		}
	}(rows)

	var episodes []*entities.EpisodeWithImage
	for rows.Next() {
		episode := &entities.EpisodeWithImage{}
		if err := rows.Scan(&episode.ID, &episode.Name, &episode.Slug,
			&episode.Season.ImageUrl); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		episodes = append(episodes, episode)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return episodes, nil
}

func (r *repository) GetByID(id uint64) (*entities.Episode, error) {
	episode := &entities.Episode{}

	err := r.db.QueryRow(queryGetById, id).Scan(&episode.ID, &episode.Name,
		&episode.Number, &episode.Duration, &episode.Url, &episode.Slug, &episode.SeasonId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("episode with id %d %v", id, ex.ErrNotFound)
			return nil, fmt.Errorf("episode with id %d %w", id, ex.ErrNotFound)
		}
		return nil, err
	}

	return episode, nil
}

func (r *repository) GetBySlug(slug string) (*entities.EpisodeWithSeasonSlug, error) {
	episode := &entities.EpisodeWithSeasonSlug{}

	err := r.db.QueryRow(queryGetBySlug, slug).Scan(&episode.ID, &episode.Name,
		&episode.Number, &episode.Duration, &episode.Url, &episode.Slug, &episode.Season.Slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("episode with slug %s %w", slug, ex.ErrNotFound)
		}
		return nil, err
	}

	return episode, nil
}

func (r *repository) Create(episode *entities.Episode) error {
	err := r.db.QueryRow(queryCreate, episode.Name, episode.Number,
		episode.Duration, episode.Url, episode.Slug, episode.SeasonId).Scan(&episode.ID)
	if err != nil {

		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code {
			case "23505": // duplicados
				return &ex.ErrAlreadyExists{
					Field:      pgErr.Column,
					Constraint: pgErr.Constraint,
				}
			case "23502": // no nulos
				return &ex.ErrValidation{
					Field:  pgErr.Column,
					Reason: "the field cannot be empty",
				}
			default:
				log.Printf("Failed creation: %s", err)
				return fmt.Errorf("failed creation: %w", err)
			}
		}

		return err
	}

	return nil
}

func (r *repository) Update(episode *entities.Episode) error {
	result, err := r.db.Exec(queryUpdate, episode.Name, episode.Number, episode.Duration, episode.Url, episode.Slug, episode.SeasonId, episode.ID)

	if err != nil {

		if pgErr, ok := err.(*pq.Error); ok {

			switch pgErr.Code {
			case "23505": // duplicados
				return &ex.ErrAlreadyExists{
					Field:      pgErr.Column,
					Constraint: pgErr.Constraint,
				}
			case "23502": // no nulos
				return &ex.ErrValidation{
					Field:  pgErr.Column,
					Reason: "the field cannot be empty",
				}
			default:
				log.Printf("Update failed: %s", err)
				return fmt.Errorf("update failed: %w", err)
			}
		}

		log.Printf("Unknown update error: %s", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ex.ErrNotFound
	}

	return nil
}

func (r *repository) Delete(id uint64) error {
	result, err := r.db.Exec(queryDeleteById, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ex.ErrNotFound
	}

	return nil
}
