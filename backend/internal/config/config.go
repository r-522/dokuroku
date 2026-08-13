package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseURL     string
	AppPassword     string
	SessionSecret   string
	AnthropicAPIKey string
	AITimeout       time.Duration
	Port            string
}

func Load() Config {
	return Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		AppPassword:     os.Getenv("APP_PASSWORD"),
		SessionSecret:   envOr("SESSION_SECRET", os.Getenv("APP_PASSWORD")),
		AnthropicAPIKey: os.Getenv("ANTHROPIC_API_KEY"),
		AITimeout:       time.Duration(envInt("AI_TIMEOUT_SECONDS", 5)) * time.Second,
		Port:            envOr("PORT", "8080"),
	}
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}
