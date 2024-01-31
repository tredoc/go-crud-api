package main

import (
	"fmt"
	"net/http"
)

type config struct {
	port int
	env  string
}

type application struct {
	cfg config
}

func main() {
	cfg := config{
		port: 3003,
		env:  "dev",
	}

	app := &application{
		cfg: cfg,
	}

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
	}

	fmt.Printf("starting %s server on %d port\n", app.cfg.env, app.cfg.port)
	err := srv.ListenAndServe()
	fmt.Println(err)
}
