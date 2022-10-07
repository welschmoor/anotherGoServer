package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string // env INT or PROD
}

type AppStatus struct {
	Status  string
	Env     string
	Version string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "Listening on port 8080")
	flag.StringVar(&cfg.env, "env", "INT", "App environmen")
	flag.Parse()

	fmt.Println("Hello B...")

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:  "Available",
			Env:     cfg.env,
			Version: version,
		}

		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)

	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}
}
