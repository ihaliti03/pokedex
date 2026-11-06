package commands

import (
	"fmt"

	"github.com/ihaliti03/pokedex/internal/cli"
	"github.com/ihaliti03/pokedex/internal/pokeapi"
)

func CommandMapB(cfg *cli.Config, _ []string) error {
	if cfg.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	page, err := pokeapi.GetLocationAreas(cfg.Cache, cfg.PreviousURL)
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