package pokeapi

type pokeapiRequester struct{}

func NewRequest() Interface {
	return &pokeapiRequester{}
}
