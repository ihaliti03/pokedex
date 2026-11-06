package commands

import (
	"fmt"

	"github.com/ihaliti03/pokedex/internal/cli"
	"github.com/ihaliti03/pokedex/internal/pokeapi"
)

func CommandMap(cfg *cli.Config, _ []string) error {
	page, err := pokeapi.GetLocationAreas(cfg.Cache, cfg.NextURL)
	if err != nil {
		return err
	}

	for _, r := range page.Results {
		fmt.Println(r.Name)
	}

	if page.Next != nil {
		cfg.NextURL = *page.Next
	} else {
		cfg.NextURL = ""
	}
	if page.Previous != nil {
		cfg.PreviousURL = *page.Previous
	} else {
		cfg.PreviousURL = ""
	}
	return nil
}