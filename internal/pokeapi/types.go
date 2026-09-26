// A pokeapi types implemented following the documentation of pokeapi endpoints
// https://pokeapi.co/docs/v2#resource-listspagination-section

package pokeapi

type ShallowList struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
