package main

import (
	"github.com/wicho90/anime-api/config"
	"github.com/wicho90/anime-api/database"
	"github.com/wicho90/anime-api/internal/episode"
	"github.com/wicho90/anime-api/internal/season"
	"github.com/wicho90/anime-api/internal/server"
	"github.com/wicho90/anime-api/internal/validator"
	"go.uber.org/fx"
	"log"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.New,
			database.New,
			season.NewRepository,
			season.NewService,
			season.NewHandler,
			episode.NewRepository,
			episode.NewService,
			episode.NewHandler,
			server.New,
			func() validator.Validator {
				return validator.NewCustomValidator()
			},
		),
		fx.Invoke(server.Start),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}
}
