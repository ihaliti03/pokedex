package commands

import (
	"fmt"

	"github.com/ihaliti03/pokedex/internal/cli"
	"github.com/ihaliti03/pokedex/internal/pokeapi"
)

func CommandExplore(cfg *cli.Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <area-name>")
	}
	area, err := pokeapi.GetLocationArea(cfg.Cache, args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}