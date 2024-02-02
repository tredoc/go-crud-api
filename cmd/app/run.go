package main

import (
	"database/sql"
	"fmt"
	"github.com/tredoc/go-crud-api/internal/handler"
	"github.com/tredoc/go-crud-api/pkg/log"
	"net/http"
)

func runDB(cfg *config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.db.host, cfg.db.port, cfg.db.username, cfg.db.password, cfg.db.name, cfg.db.sslmode)

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

func runServer(cfg *config, handlers *handler.Handler) error {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: handlers.InitRoutes(),
	}

	log.Info(fmt.Sprintf("starting %s server on %s port", cfg.env, cfg.port))
	return srv.ListenAndServe()
}
