package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationInfo struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationInfo(name string) (LocationInfo, error) {
	url := ApiURL + "/location-area/" + name
	data, ok := c.Cache.Get(url)
	if ok {
		return readLocationInfo(data)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationInfo{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, nil
	}

	c.Cache.Add(url, data)
	areas, err := readLocationInfo(data)
	if err != nil {
		return LocationInfo{}, err
	}

	return areas, nil
}

func readLocationInfo(data []byte) (LocationInfo, error) {
	info := LocationInfo{}
	if err := json.Unmarshal(data, &info); err != nil {
		return LocationInfo{}, fmt.Errorf("Could not read locations: %v", err)
	}
	return info, nil
}
