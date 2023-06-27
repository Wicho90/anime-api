package episode

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wicho90/anime-api/internal/entities"
)

type Repository interface {
	GetAll() ([]*entities.Episode, error)
	GetLatest() ([]*entities.EpisodeWithImage, error)
	GetByID(id uint64) (*entities.Episode, error)
	GetBySlug(slug string) (*entities.EpisodeWithSeasonSlug, error)
	Create(episode *entities.Episode) error
	Update(episode *entities.Episode) error
	Delete(id uint64) error
}

type Service interface {
	GetAll() ([]*entities.Episode, error)
	GetLatest() ([]*entities.EpisodeWithImage, error)
	GetByID(id uint64) (*entities.Episode, error)
	GetBySlug(slug string) (*entities.EpisodeWithSeasonSlug, error)
	Create(episode *entities.Episode) (*entities.Episode, error)
	Update(id uint64, episode *entities.Episode) (*entities.Episode, error)
	Delete(id uint64) error
}

type Handler interface {
	GetAll(ctx *fiber.Ctx) error
	GetLatest(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetBySlug(ctx *fiber.Ctx) error
}
