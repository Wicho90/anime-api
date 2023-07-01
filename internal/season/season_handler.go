package season

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/wicho90/anime-api/internal/entities"
	"github.com/wicho90/anime-api/internal/ex"
	"github.com/wicho90/anime-api/internal/response"
	"github.com/wicho90/anime-api/internal/validator"
	"net/http"
	"strconv"
)

type handler struct {
	service   Service
	validator validator.Validator
}

func NewHandler(service Service, validator validator.Validator) Handler {
	return &handler{
		service:   service,
		validator: validator,
	}
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	seasons, err := h.service.GetAll()
	if err != nil {
		return response.NewInternalServerErrorResponse("Failed to get seasons")
	}

	return ctx.Status(http.StatusOK).JSON(seasons)
}

func (h *handler) GetById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.NewBadRequestResponse("Invalid id")
	}

	season, err := h.service.GetById(id)
	if err != nil {

		if errors.Is(err, ex.ErrNotFound) {
			return response.NewNotFoundResponse(err.Error())
		}

		return response.NewInternalServerErrorResponse("Failed to get season")
	}

	return ctx.Status(http.StatusOK).JSON(season)

}

func (h *handler) Create(ctx *fiber.Ctx) error {
	var season *entities.Season

	if err := ctx.BodyParser(&season); err != nil {
		return response.NewBadRequestResponse("Invalid request body")
	}

	if m, err := h.validator.Validate(season); err != nil {
		return response.NewBadRequestResponse(m)
	}

	err := h.service.Create(season)
	if err != nil {
		return response.NewInternalServerErrorResponse("Failed to create episode")
	}

	return ctx.Status(http.StatusCreated).
		JSON(season)
}

func (h *handler) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.NewBadRequestResponse("Invalid id")
	}

	var season *entities.Season
	if err := ctx.BodyParser(&season); err != nil {
		return response.NewBadRequestResponse("Invalid request body")
	}

	if m, err := h.validator.Validate(season); err != nil {
		return response.NewBadRequestResponse(m)
	}

	err = h.service.Update(id, season)
	if err != nil {
		if errors.Is(err, ex.ErrNotFound) {
			return response.NewNotFoundResponse(err.Error())
		}

		return response.NewInternalServerErrorResponse("Failed to update season")
	}

	return ctx.Status(http.StatusOK).JSON(season)
}

func (h *handler) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.NewBadRequestResponse("Invalid id")
	}

	err = h.service.Delete(id)
	if err != nil {

		if errors.Is(err, ex.ErrNotFound) {
			return response.NewNotFoundResponse(err.Error())
		}

		return response.NewInternalServerErrorResponse("Failed to delete season")
	}

	return ctx.SendStatus(http.StatusNoContent)
}
