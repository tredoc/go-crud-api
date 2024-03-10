package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
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

func runCache(cfg *config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.cache.host, cfg.cache.port),
		Password: cfg.cache.password,
		DB:       cfg.cache.dbs,
	})

	res := rdb.Ping(context.Background())
	if err := res.Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}

func runServer(cfg *config, handlers *handler.Handler) error {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: handlers.InitRoutes(),
	}

	log.Info(fmt.Sprintf("starting %s server on %s port", cfg.env, cfg.port))
	return srv.ListenAndServe()
}
