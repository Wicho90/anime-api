package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wicho90/anime-api/config"
)

func New(config *config.Config) (*sql.DB, error) {
	host := config.Database.Host
	port := config.Database.Port
	user := config.Database.User
	password := config.Database.Password
	dbName := config.Database.Name

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
