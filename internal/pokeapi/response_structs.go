package pokeapi

type ResponseLocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type ResponseListPokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type ResponsePokemonInfo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"` 
}