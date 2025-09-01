package client

import (
	"net/http"
	"time"

	"pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{},
		cache: pokecache.NewCache(4 * time.Second),
	}
}
