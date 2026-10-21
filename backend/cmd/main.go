package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SemgaTeam/hackaton2026_autumn/backend/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("config: ", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	log.Printf("listening on %s", cfg.Addr())
	log.Fatal(http.ListenAndServe(cfg.Addr(), mux))
}
