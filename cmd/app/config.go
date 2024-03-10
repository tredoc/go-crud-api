package main

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type dbConfig struct {
	host     string
	username string
	password string
	name     string
	port     string
	sslmode  string
}

type cacheConfig struct {
	host     string
	port     string
	password string
	dbs      int
}

type config struct {
	port  string
	env   string
	db    dbConfig
	cache cacheConfig
}

func getConfig() (*config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	redisDBS, err := strconv.Atoi(os.Getenv("REDIS_DBS"))
	if err != nil {
		return nil, errors.New("can't convert redis port to int")
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
		cache: cacheConfig{
			host:     os.Getenv("REDIS_HOST"),
			port:     os.Getenv("REDIS_PORT"),
			password: os.Getenv("REDIS_PASSWORD"),
			dbs:      redisDBS,
		},
	}, nil
}
