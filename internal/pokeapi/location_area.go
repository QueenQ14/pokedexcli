package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(pageURL *string) (ResponseLocationArea, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		var res ResponseLocationArea
		if err := json.Unmarshal(val, &res); err != nil {
			return ResponseLocationArea{}, err
		}
		return res, nil
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

	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationArea{}, err
	}
	c.cache.Add(url, rawData)

	var data ResponseLocationArea
	if err := json.Unmarshal(rawData, &data); err != nil {
		return ResponseLocationArea{}, err
	}
	return data, nil
}
