package pokeapi

import "fmt"

type PokeResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (pr *PokeResponse) Print() {
	for _, r := range pr.Results {
		fmt.Println(r.Name)
	}
}
