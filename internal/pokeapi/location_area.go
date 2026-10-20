package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocationArea(pageURL *string) (ResponseLocationArea, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}
	rawData, err := c.get(url)
	if err != nil {
		return ResponseLocationArea{}, err
	}
	var data ResponseLocationArea
	if err := json.Unmarshal(rawData, &data); err != nil {
		return ResponseLocationArea{}, err
	}
	return data, nil
}
