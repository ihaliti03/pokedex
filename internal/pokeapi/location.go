package pokeapi

import (
	"fmt"

	"github.com/ihaliti03/pokedex/internal/pokecache"
)

type LocationAreaPage struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []LocationEntry `json:"results"`
}

type LocationEntry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

// GetLocationAreas fetches a page of location areas.
// url may be the base URL or a next/previous link from a prior response.
func GetLocationAreas(cache *pokecache.Cache, url string) (*LocationAreaPage, error) {
	if url == "" {
		url = baseURL + "/location-area/"
	}
	return fetchCached[LocationAreaPage](cache, url)
}

// GetLocationArea fetches a single location area by name.
func GetLocationArea(cache *pokecache.Cache, name string) (*LocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s/", baseURL, name)
	return fetchCached[LocationArea](cache, url)
}