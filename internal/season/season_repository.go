package season

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/wicho90/anime-api/internal/entities"
	"github.com/wicho90/anime-api/internal/ex"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	queryGetAll  = "SELECT id, name, number, slug, image_url FROM seasons"
	queryGetById = "SELECT id, name, number, slug, image_url FROM seasons WHERE id = $1"
	queryCreate  = "INSERT INTO seasons (name, number, slug, image_url) VALUES ($1, $2, $3, $4) RETURNING id"
	queryUpdate  = "UPDATE seasons SET name = $1, number = $2, slug = $3, image_url = $4 WHERE id = $5"

	queryDeleteById = "DELETE FROM seasons WHERE id = $1"
)

func (r *repository) GetAll() ([]*entities.Season, error) {
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

	var seasons []*entities.Season
	for rows.Next() {
		season := &entities.Season{}
		if err := rows.Scan(&season.ID, &season.Name, &season.Number, &season.Slug, &season.ImageUrl); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		seasons = append(seasons, season)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return seasons, nil
}

func (r *repository) GetById(id uint64) (*entities.Season, error) {
	season := &entities.Season{}

	err := r.db.QueryRow(queryGetById, id).
		Scan(&season.ID, &season.Name, &season.Number, &season.Slug, &season.ImageUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("season with id %d %v", id, ex.ErrNotFound)
			return nil, fmt.Errorf("season with id %d %w", id, ex.ErrNotFound)
		}
		return nil, err
	}

	return season, nil
}

func (r *repository) Create(season *entities.Season) error {
	err := r.db.QueryRow(queryCreate, season.Name,
		season.Number, season.Slug, season.ImageUrl).Scan(&season.ID)

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

func (r *repository) Update(season *entities.Season) error {
	result, err := r.db.Exec(queryUpdate, season.Name, season.Number, season.Slug, season.ImageUrl, season.ID)

	if err != nil {

		if pgErr, ok := err.(*pq.Error); ok {

			switch pgErr.Code {
			case "23505":
				return &ex.ErrAlreadyExists{
					Field:      pgErr.Column,
					Constraint: pgErr.Constraint,
				}
			case "23502":
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
