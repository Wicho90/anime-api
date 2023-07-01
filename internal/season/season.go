package season

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wicho90/anime-api/internal/entities"
)

type Repository interface {
	GetAll() ([]*entities.Season, error)
	GetById(id uint64) (*entities.Season, error)
	Create(season *entities.Season) error
	Update(season *entities.Season) error
	Delete(id uint64) error
}

type Service interface {
	GetAll() ([]*entities.Season, error)
	GetById(id uint64) (*entities.Season, error)
	Create(season *entities.Season) error
	Update(id uint64, season *entities.Season) error
	Delete(id uint64) error
}

type Handler interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
