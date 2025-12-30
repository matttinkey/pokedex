package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const resultsPerPage int = 20

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationsUrl(numResults, page int) string {
	offset := page * numResults
	url := ApiURL + "/location-area" + fmt.Sprintf("?limit=%v&offset=%v", numResults, offset)
	return url
}

func (c *Client) GetLocationAreas(url string) (LocationAreas, error) {
	data, ok := c.Cache.Get(url)
	if ok {
		return readLocationsData(data)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, nil
	}

	c.Cache.Add(url, data)
	areas, err := readLocationsData(data)
	if err != nil {
		return LocationAreas{}, err
	}

	return areas, nil
}

func readLocationsData(data []byte) (LocationAreas, error) {
	areas := LocationAreas{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return LocationAreas{}, fmt.Errorf("Could not read locations: %v", err)
	}
	return areas, nil
}
