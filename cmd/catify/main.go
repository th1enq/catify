package main

import (
	"catify/internal/cache"
	"catify/internal/config"
	"catify/internal/database"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: %v", err)
	}

	db, err := database.InitPostgres(struct {
		Port     string
		Host     string
		User     string
		Password string
		Name     string
	}(cfg.Database))
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: %v", err)
	}

	es, err := database.InitElasticsearch(struct{ URL string }(cfg.Elasticsearch))
	if err != nil {
		log.Fatal("Failed to connect to Elasticsearch: %v", err)
	}

	redisCache, err := cache.NewRedisCache(
		fmt.Sprintf("%s %s", cfg.Redis.Host, cfg.Redis.Port),
		cfg.Redis.Password,
		cfg.Redis.DB,
	)
	if err != nil {
		log.Fatal("Failed to connect to Redis: %v", err)
	}


	
}
