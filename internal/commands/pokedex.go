package commands

import (
	"fmt"
	"sort"

	"github.com/ihaliti03/pokedex/internal/cli"
)

func CommandPokedex(cfg *cli.Config, _ []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")

	names := make([]string, 0, len(cfg.Pokedex))
	for name := range cfg.Pokedex {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}