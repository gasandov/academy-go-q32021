package entities

type APIResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []PokemonAPI `json:"results"`
}

type PokemonAPI struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
}
