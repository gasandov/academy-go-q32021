package entities

type API struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []PokemonResponse `json:"results"`
}
