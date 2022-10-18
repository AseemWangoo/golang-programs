package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	client *redis.Client
}

func NewRedis() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		DB:          0, // use default DB
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}
