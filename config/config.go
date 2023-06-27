package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

func New() *Config {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env. Se utilizarán las configuraciones predeterminadas.")
	}

	cfg.Server.Port = os.Getenv("SERVER_PORT")
	cfg.Database.Host = os.Getenv("DB_HOST")
	cfg.Database.Port = os.Getenv("DB_PORT")
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.Name = os.Getenv("DB_NAME")

	cfg.standard()

	log.Printf("Config: %v", cfg)

	return cfg
}

func (c *Config) standard() {
	if c.Server.Port == "" {
		c.Server.Port = "8080"
	}
	if c.Database.Host == "" {
		c.Database.Host = "localhost"
	}
	if c.Database.Port == "" {
		c.Database.Port = "5432"
	}
	if c.Database.User == "" {
		c.Database.User = "postgres"
	}
	if c.Database.Password == "" {
		c.Database.Password = "sasa"
	}
	if c.Database.Name == "" {
		c.Database.Name = "animedb"
	}
}
