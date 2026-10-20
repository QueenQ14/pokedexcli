package pokeapi

import "encoding/json"

func (c *Client) GetPokemonInfo(pkmn string) (ResponsePokemonInfo, error) {
	url := baseURL + "/pokemon/" + pkmn
	rawData, err := c.get(url)
	if err != nil {
		return ResponsePokemonInfo{}, err
	}
	var data ResponsePokemonInfo
	if err := json.Unmarshal(rawData, &data); err != nil {
		return ResponsePokemonInfo{}, err
	}
	return data, nil

}
