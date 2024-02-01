package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tredoc/go-crud-api/internal/handler"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/internal/service"
	"log"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type dbConfig struct {
	host     string
	username string
	password string
	name     string
	port     string
	sslmode  string
}

type config struct {
	port string
	env  string
	db   dbConfig
}

type application struct {
	cfg     config
	log     *slog.Logger
	handler *handler.Handler
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		port: os.Getenv("PORT"),
		env:  os.Getenv("ENV"),
		db: dbConfig{
			host:     os.Getenv("DB_HOST"),
			username: os.Getenv("DB_USERNAME"),
			password: os.Getenv("DB_PASSWORD"),
			name:     os.Getenv("DB_NAME"),
			port:     os.Getenv("DB_PORT"),
			sslmode:  os.Getenv("SSL_MODE"),
		},
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.db.host, cfg.db.port, cfg.db.username, cfg.db.password, cfg.db.name, cfg.db.sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := &application{
		cfg: cfg,
		log: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: handlers.InitRoutes(),
	}

	app.log.Info(fmt.Sprintf("starting %s server on %s port", app.cfg.env, app.cfg.port))
	err = srv.ListenAndServe()
	app.log.Error(err.Error())
}
