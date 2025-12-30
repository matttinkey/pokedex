package pokeapi

import (
	"net/http"
	"pokedexcli/internal/cache"
	"time"
)

const ApiURL string = "https://pokeapi.co/api/v2"
const defaultTimeout = time.Minute

type Client struct {
	httpClient http.Client
	Cache      cache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: defaultTimeout,
		},
		Cache: cache.NewCache(cacheInterval),
	}
}
