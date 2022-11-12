package pokeapi

type PokeapiPaginationResponseDTO struct {
	Count   uint64                               `json:"count"`
	Results []PokeapiPaginationResponseResultDTO `json:"results"`
}

type PokeapiPaginationResponseResultDTO struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokeapiGetByIDResponseDTO struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Height uint64 `json:"height"`
	Weight uint64 `json:"weight"`
}
