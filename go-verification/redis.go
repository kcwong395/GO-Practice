package main

import (
	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
}

func Init() *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Client{rdb}
}
