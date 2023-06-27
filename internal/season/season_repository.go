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
					Reason: "no puede estar vacío",
				}
			default:
				log.Printf("Error de restricción en la actualización: %s", err)
				return fmt.Errorf("error en la creación: %w", err)
			}
		}

		return err
	}

	return nil
}

func (r *repository) Update(id uint64, season entities.Season) (entities.Season, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}
