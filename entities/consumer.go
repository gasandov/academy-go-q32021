package entities

type APIResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []APIPokemon `json:"results"`
}
