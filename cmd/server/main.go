package main

import (
	"catify/internal/config"
	"catify/internal/db"
	"catify/internal/redis"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: %v", err)
	}

	cfgDB, err := db.Init(cfg)
	if err != nil {
		log.Fatal("Failed to load database configuration: %v", err)
	}

	cfgRedis, err := redis.Init(cfg)
	if err != nil {
		log.Fatal("Failed to load redis configuration: %v", err)
	}
}
