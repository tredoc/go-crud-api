package main

import (
	_ "github.com/lib/pq"
	"github.com/tredoc/go-crud-api/internal/handler"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/internal/service"
	"github.com/tredoc/go-crud-api/pkg/log"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := runDB(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	err = runServer(cfg, handlers)
	if err != nil {
		log.Fatal(err.Error())
	}
}
