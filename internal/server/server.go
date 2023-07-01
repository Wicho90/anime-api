package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wicho90/anime-api/config"
	"github.com/wicho90/anime-api/internal/episode"
	"github.com/wicho90/anime-api/internal/response"
	"github.com/wicho90/anime-api/internal/season"
	"log"
	"net/http"
)

type Server struct {
	app *fiber.App
}

func errorHandler(ctx *fiber.Ctx, err error) error {

	if seer, ok := err.(response.BaseResponse); ok {
		return ctx.Status(seer.GetCode()).
			JSON(seer)
	}

	return ctx.Status(http.StatusInternalServerError).
		JSON(err)
}

func New(
	seasonHandler season.Handler,
	episodeHandler episode.Handler,
) *Server {
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5173",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type",
		AllowCredentials: false,
	}))

	app.Static("/", "./public")

	v1 := app.Group("/api/v1")
	{
		seasons := v1.Group("seasons")
		{
			seasons.Get("/", seasonHandler.GetAll)
			seasons.Get("/:id", seasonHandler.GetById)
			seasons.Post("/", seasonHandler.Create)
			seasons.Put("/:id", seasonHandler.Update)
			seasons.Delete("/:id", seasonHandler.Delete)
		}
		episodes := v1.Group("/episodes")
		{
			episodes.Get("/", episodeHandler.GetAll)
			episodes.Get("/latest", episodeHandler.GetLatest)
			episodes.Get("/:id", episodeHandler.GetById)
			episodes.Get("/slug/:slug", episodeHandler.GetBySlug)
			episodes.Post("/", episodeHandler.Create)
			episodes.Put("/:id", episodeHandler.Update)
			episodes.Delete("/:id", episodeHandler.Delete)
		}
	}

	return &Server{
		app: app,
	}
}
func Start(s *Server, config *config.Config) {
	addr := fmt.Sprintf(":%s", config.Server.Port)
	log.Printf("Server listening on %s", addr)

	err := s.app.Listen(addr)
	if err != nil {
		panic(err)
	}
}
