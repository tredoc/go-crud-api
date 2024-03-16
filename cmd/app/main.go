package main

import (
	_ "github.com/lib/pq"
	"github.com/tredoc/go-crud-api/internal/cache"
	"github.com/tredoc/go-crud-api/internal/handler"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/internal/service"
	"github.com/tredoc/go-crud-api/pkg/log"
)

// @title Swagger go-crud-api API
// @version 4.0
// @description personal educational project

// @host localhost:3000
// @BasePath /

// @SecurityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rdb, err := runRCache(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rdb.Close()

	db, err := runDB(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	rch := cache.NewCache(rdb)
	repos := repository.NewRepository(db)
	services := service.NewService(repos, rch)
	handlers := handler.NewHandler(services)

	err = runServer(cfg, handlers)
	if err != nil {
		log.Fatal(err.Error())
	}
}
