package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"sheltes-map/internal/config"
	"sheltes-map/internal/repository"
	"sheltes-map/internal/service"
	"sheltes-map/internal/transport/grpc"
	"sheltes-map/pkg/database"
)

func main() {
	log := logrus.New()

	conf, err := config.NewConfig()
	if err != nil {
		log.Error("error while init config", err)
	}

	db, err := database.NewDatabase(conf.Database)
	if err != nil {
		log.Fatalf("error while run database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	serv := grpc.NewGrpcServer(services)

	if err := serv.ListenAndServe(conf.Server); err != nil {
		log.Fatalf("error while run server: %s", err.Error())
	}
}
