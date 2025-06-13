package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "location"
	fullURL := baseURL + endpoint

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

	return locationAreasRep, nil

}
