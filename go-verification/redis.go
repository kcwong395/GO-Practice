package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
}

func Init() (*Client, error) {
	// Define the connection configuration
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ensure the connection is successful
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return &Client{rdb}, nil
}
