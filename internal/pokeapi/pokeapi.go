package pokeapi

import (
	"net/http"
	"time"

	"github.com/dakotatokad/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval), // Cache with an interval
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
