package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonArea(area string) (ResponseListPokemon, error) {
	url := baseURL + "/location-area/" + area
	if val, ok := c.cache.Get(url); ok {
		var res ResponseListPokemon
		if err := json.Unmarshal(val, &res); err != nil {
			return ResponseListPokemon{}, err
		}
		return res, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseListPokemon{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseListPokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return ResponseListPokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseListPokemon{}, err
	}
	c.cache.Add(url, rawData)

	var data ResponseListPokemon
	if err := json.Unmarshal(rawData, &data); err != nil {
		return ResponseListPokemon{}, err
	}
	return data, nil
}
