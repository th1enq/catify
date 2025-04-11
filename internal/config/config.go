package config

import (
	"fmt"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT,required"`

	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,requried"`
	DBPassword string `env:"DB_PASSWORD,requried"`
	DBName     string `env:"DB_NAME,requried"`

	RedisHost     string `env:"REDIS_HOST,requried"`
	RedisPort     int    `env:"REDIS_PORT,requried"`
	RedisPassword string `env:"REDIS_PASSWORD,requried"`
	RedisDB       int    `env:"REDIS_DB,requried"`

	ElasticPassword string `env:"ELASTIC_PASSWORD,required"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config

	if err := envdecode.Decode(&cfg); err != nil {
		return nil, err
	}

	fmt.Println(cfg)

	return &cfg, nil
}
