package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sheltes-map/internal/config"
)

func NewDatabase(env *config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.Host, env.Port, env.User, env.Password, env.Name, env.SslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
