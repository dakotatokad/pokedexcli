package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// Check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// hit cache
		log.Print("Cache hit for URL: ", fullURL)
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return pokemon, nil
	}
	log.Print("Cache miss for URL: ", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("failed to fetch Pokemon: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to read response body: %w", err)
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil

}
