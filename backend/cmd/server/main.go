package main

import (
	"log"
	"urlshortener/internal/config"
	"urlshortener/internal/db"
	"urlshortener/internal/cache"
)

func main(){
	cfg := config.Load()

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set. Check your environment / .env file.")
	}

	database, err := db.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	defer database.Close()
	log.Println("Connected to Postgres")

	redisCache, err := cache.New(cfg.RedisURL)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	defer redisCache.Close()
	log.Println("Connected to Redis")


}