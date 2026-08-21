package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)


type Cache struct client *redis.Client


func New(redisURL string) (*Cache, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil return nil, err
	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil return nil, err
	return &Cache{client: client}, nil
}

func (c *Cache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}