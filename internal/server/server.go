package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wicho90/anime-api/config"
	"github.com/wicho90/anime-api/internal/episode"
	"github.com/wicho90/anime-api/internal/season"
	"log"
)

type Server struct {
	app *fiber.App
}

func New(
	seasonHandler season.Handler,
	episodeHandler episode.Handler,
) *Server {
	app := fiber.New()
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
			/*seasons.Put("/:id")
			seasons.Delete("/:id")*/
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
