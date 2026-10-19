package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetLocationArea(pageURL *string) (ResponseLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocationArea{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocationArea{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return ResponseLocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	var data ResponseLocationArea
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return ResponseLocationArea{}, err
	}
	return data, nil
}
