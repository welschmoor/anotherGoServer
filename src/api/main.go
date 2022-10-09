package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string // env INT or PROD
}

type AppStatus struct {
	Status  string `json:"status"` //in json make the field be named status (lowercase)
	Env     string `json:"env"`
	Version string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "Listening on port 8080")
	flag.StringVar(&cfg.env, "env", "INT", "App environmen")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // | means pipe

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Running on port", cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
