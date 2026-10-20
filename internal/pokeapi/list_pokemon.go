package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListPokemonArea(area string) (ResponseListPokemon, error) {
	url := baseURL + "/location-area/" + area
	rawData, err := c.get(url)
	if err != nil {
		return ResponseListPokemon{}, err
	}
	var data ResponseListPokemon
	if err := json.Unmarshal(rawData, &data); err != nil {
		return ResponseListPokemon{}, err
	}
	return data, nil
}
