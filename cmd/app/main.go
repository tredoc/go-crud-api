package main

import (
	"fmt"
	"github.com/tredoc/go-crud-api/internal/handler"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	port int
	env  string
}

type application struct {
	cfg     config
	log     *slog.Logger
	handler *handler.Handler
}

func main() {
	cfg := config{
		port: 3003,
		env:  "dev",
	}

	app := &application{
		cfg:     cfg,
		log:     slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		handler: handler.NewHandler(),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.handler.InitRoutes(),
	}

	app.log.Info(fmt.Sprintf("starting %s server on %d port", app.cfg.env, app.cfg.port))
	err := srv.ListenAndServe()
	app.log.Error(err.Error())
}
