package pokeapi

import (
	"fmt"

	"github.com/ihaliti03/pokedex/internal/pokecache"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

// GetPokemon fetches a Pokemon by name.
func GetPokemon(cache *pokecache.Cache, name string) (*Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s/", baseURL, name)
	return fetchCached[Pokemon](cache, url)
}