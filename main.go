package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"sheltes-map/internal/config"
	"sheltes-map/internal/importer"
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

	repos := repository.NewRepository(db)
	imp := importer.NewImporter(repos.Shelter)

	if err := imp.Import(); err != nil {
		fmt.Println(err)
	}

}
