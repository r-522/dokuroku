package main

import (
	"database/sql"
	"log"
	"net/http"

	"dokuroku/backend/internal/ai"
	"dokuroku/backend/internal/auth"
	"dokuroku/backend/internal/config"
	httpapi "dokuroku/backend/internal/http"
	"dokuroku/backend/internal/note"
)

func main() {
	cfg := config.Load()
	if cfg.AppPassword == "" || cfg.DatabaseURL == "" {
		log.Fatal("APP_PASSWORD and DATABASE_URL are required")
	}
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	server := httpapi.Server{
		Auth: auth.NewManager(cfg.AppPassword, cfg.SessionSecret),
		Service: note.Service{
			Repo: note.SQLRepository{DB: db},
			AI:   ai.ClaudeClient{APIKey: cfg.AnthropicAPIKey, Timeout: cfg.AITimeout},
		},
	}
	log.Printf("listening on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, server.Routes()))
}
