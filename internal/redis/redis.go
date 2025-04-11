package redis

import (
	"catify/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
}

func Init(cfg *config.Config) (*Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}
