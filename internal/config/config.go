package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server configuration
	ServerPort string

	// Database configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Redis configuration
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		ServerPort:    getEnvOrDefaultString("SERVER_PORT", "8080"),
		DBHost:        getEnvOrDefaultString("DB_HOST", "postgres"),
		DBPort:        getEnvOrDefaultString("DB_PORT", "5432"),
		DBUser:        getEnvOrDefaultString("DB_USER", "mydb"),
		DBPassword:    getEnvOrDefaultString("DB_PASSWORD", "admin"),
		DBName:        getEnvOrDefaultString("DB_NAME", "catify"),
		DBSSLMode:     getEnvOrDefaultString("DB_SSLMODE", "disable"),
		RedisHost:     getEnvOrDefaultString("REDIS_HOST", "localhost"),
		RedisPort:     getEnvOrDefaultString("REDIS_PORT", "6379"),
		RedisPassword: getEnvOrDefaultString("REDIS_PASSWORD", "password"),
		RedisDB:       getEnvOrDefaultInt("REDIS_DB", 0),
	}, nil
}

func getEnvOrDefaultString(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist != false {
		return value
	}
	return defaultValue
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	if value, exist := os.LookupEnv(key); exist != false {
		IntValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return IntValue
	}
	return defaultValue
}
