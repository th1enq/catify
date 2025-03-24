package cache

import (
	"catify/internal/models"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string, password string, db int) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{
		client: client,
	}, nil
}

func (r *RedisCache) SetSong(ctx context.Context, song *models.Song) error {
	songJSON, err := json.Marshal(song)
	if err != nil {
		return err
	}
	key := "song:" + strconv.FormatUint(uint64(song.ID), 10)

	return r.client.Set(ctx, key, songJSON, 24*time.Hour).Err()
}

func (r *RedisCache) GetSong(ctx context.Context, id uint) (*models.Song, error) {
	key := "song:" + strconv.FormatUint(uint64(id), 10)

	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var song models.Song
	err = json.Unmarshal([]byte(val), &song)

	return &song, err
}

func (r *RedisCache) SetSearchResults(ctx context.Context, query string, songs []models.Song) error {
	songJSON, err := json.Marshal(songs)
	if err != nil {
		return err
	}

	key := "search:" + query
	return r.client.Set(ctx, key, songJSON, 24*time.Hour).Err()
}

func (r *RedisCache) GetSearchResults(ctx context.Context, query string) ([]models.Song, error) {
	key := "search:" + query

	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var songs []models.Song
	err = json.Unmarshal([]byte(val), &songs)

	return songs, err
}
