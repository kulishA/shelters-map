package main

import (
	"sheltes-map/internal/config"
	"sheltes-map/internal/logger"
	"sheltes-map/internal/repository"
)

func main() {
	log := logger.NewLogger()

	conf, err := config.NewConfig()
	if err != nil {
		log.Error("error while init config", err)
	}

	db, err := repository.NewDatabase(conf.Database)
	if err != nil {
		log.Error("error while run database", err)
	}

}
