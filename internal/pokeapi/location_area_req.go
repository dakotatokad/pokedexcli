package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// Check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// hit cache
		log.Print("Cache hit for URL: ", fullURL)
		locationAreasRep := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasRep)
		if err != nil {
			return LocationAreasResp{}, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return locationAreasRep, nil
	}
	log.Print("Cache miss for URL: ", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("failed to fetch location areas: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("failed to read response body: %w", err)
	}

	locationAreasRep := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasRep)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.cache.Add(fullURL, data)

	return locationAreasRep, nil

}
