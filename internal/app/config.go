package app

import (
	"os"
	"strconv"
	"time"
)

// Config holds all application configuration
type Config struct {
	Port            int
	JWTSecret       string
	DatabaseURL     string
	IdleTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	RateLimit       int
}

// DefaultConfig returns the default configuration loaded from environment
func DefaultConfig() *Config {
	return &Config{
		Port:            getEnvInt("PORT", 8080),
		JWTSecret:       getEnvString("JWT_SECRET", "some-secret-123456"),
		DatabaseURL:     getEnvString("DATABASE_URL", ""),
		IdleTimeout:     time.Minute,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    30 * time.Second,
		ShutdownTimeout: 5 * time.Second,
		RateLimit:       100,
	}
}

func getEnvString(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
