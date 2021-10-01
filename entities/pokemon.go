package entities

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type APIPokemon struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
