package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sheltes-map/internal/config"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase(env *config.DatabaseConfig) (*Database, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.Host, env.Port, env.User, env.Password, env.Name, env.SslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	database := Database{DB: db}
	return &database, nil
}

func (d *Database) ShutDown() error {
	return d.DB.Close()
}
