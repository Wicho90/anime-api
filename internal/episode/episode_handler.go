package episode

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/wicho90/anime-api/internal/entities"
	"github.com/wicho90/anime-api/internal/ex"
	"github.com/wicho90/anime-api/internal/response"
	"github.com/wicho90/anime-api/internal/validator"
	"log"
	"net/http"
	"strconv"
)

type handler struct {
	service   Service
	validator validator.Validator
}

func NewHandler(service Service, validator validator.Validator) Handler {
	return &handler{service: service, validator: validator}
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	episodes, err := h.service.GetAll()
	if err != nil {
		return response.NewInternalServerErrorResponse("Failed to get episodes")
	}

	return ctx.Status(http.StatusOK).JSON(episodes)
}

func (h *handler) GetLatest(ctx *fiber.Ctx) error {
	episodes, err := h.service.GetLatest()

	if err != nil {
		log.Println(err)
		return response.NewInternalServerErrorResponse("Failed to get episodes")
	}

	return ctx.Status(http.StatusOK).JSON(episodes)
}

func (h *handler) GetById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.NewBadRequestResponse("Invalid id")
	}

	episode, err := h.service.GetByID(id)
	if err != nil {

		if errors.Is(err, ex.ErrNotFound) {
			return response.NewNotFoundResponse(err.Error())
		}

		return response.NewInternalServerErrorResponse("Failed to get episode")
	}

	return ctx.Status(http.StatusOK).JSON(episode)
}

func (h *handler) GetBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	episode, err := h.service.GetBySlug(slug)
	if err != nil {

		if errors.Is(err, ex.ErrNotFound) {
			return response.NewNotFoundResponse(err.Error())
		}

		return response.NewInternalServerErrorResponse("Failed to get episode")
	}

	return ctx.Status(http.StatusOK).JSON(episode)
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	var episode *entities.Episode

	if err := ctx.BodyParser(&episode); err != nil {
		return response.NewBadRequestResponse("Invalid request body")
	}

	if m, err := h.validator.Validate(episode); err != nil {
		return response.NewBadRequestResponse(m)
	}

	err := h.service.Create(episode)
	if err != nil {

		if errors.Is(err, ex.ErrNotFound) {
			return response.NewNotFoundResponse(err.Error())
		}

		return response.NewInternalServerErrorResponse("Failed to create episode")
	}

	return ctx.Status(http.StatusCreated).JSON(episode)
}

func (h *handler) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.NewBadRequestResponse("Invalid id")
	}

	var episode *entities.Episode
	if err = ctx.BodyParser(&episode); err != nil {
		return response.NewBadRequestResponse("Invalid request body")
	}

	if m, err := h.validator.Validate(episode); err != nil {
		return response.NewBadRequestResponse(m)
	}

	err = h.service.Update(id, episode)
	if err != nil {
		if errors.Is(err, ex.ErrNotFound) {
			return response.NewNotFoundResponse(err.Error())
		}

		return response.NewInternalServerErrorResponse("Failed to update episode")
	}

	return ctx.Status(http.StatusOK).JSON(episode)
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

		return response.NewInternalServerErrorResponse("Failed to delete episode")
	}

	return ctx.SendStatus(http.StatusNoContent)
}
