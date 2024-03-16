package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/tredoc/go-crud-api/internal/handler"
	"github.com/tredoc/go-crud-api/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

func runRCache(cfg *config) (*redis.Client, error) {
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

	go func() {
		log.Info(fmt.Sprintf("starting %s server on %s port", cfg.env, cfg.port))
		err := srv.ListenAndServe()
		if err != nil {
			log.Error("starting server error:", "error", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server Shutdown Failed: ", err.Error())
	}

	return errors.New("server gracefully stopped")
}
