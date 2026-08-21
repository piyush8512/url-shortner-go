package config

import (
	"os"
	"strconv"
)

type config struct {
	Port          string
	DatabaseURL   string
	RedisURL      string
	JWTSecret     string
	JWTExpiryDays int
	BaseURL       string
	LinkCacheTTL  int 
	Env           string

}


func Load() Config{
	return Config{
		Port:          getEnv("PORT", "3000"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://admin:changeme@localhost:5432/urlshortener?sslmode=disable"),
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:     getEnv("JWT_SECRET", ""),
		JWTExpiryDays: getEnvInt("JWT_EXPIRY_DAYS", 7),
		BaseURL:       getEnv("BASE_URL", "http://localhost:3000"),
		LinkCacheTTL:  getEnvInt("LINK_CACHE_TTL", 86400),
		Env:           getEnv("NODE_ENV", "development"),
	}
}



func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}


func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}