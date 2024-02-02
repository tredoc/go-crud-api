package main

import (
	"github.com/joho/godotenv"
	"os"
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

func getConfig() (*config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &config{
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
	}, nil
}
