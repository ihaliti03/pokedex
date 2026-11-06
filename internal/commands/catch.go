package commands

import (
	"fmt"
	"math/rand"

	"github.com/ihaliti03/pokedex/internal/cli"
	"github.com/ihaliti03/pokedex/internal/pokeapi"
)

func CommandCatch(cfg *cli.Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}
	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := pokeapi.GetPokemon(cfg.Cache, name)
	if err != nil {
		return err
	}

	chance := 50 - pokemon.BaseExperience/2
	if chance < 1 {
		chance = 1
	}

	if rand.Intn(100) < chance {
		fmt.Printf("%s was caught!\n", name)
		cfg.Pokedex[name] = *pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}