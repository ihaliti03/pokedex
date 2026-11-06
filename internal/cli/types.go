package cli

import (
	"github.com/ihaliti03/pokedex/internal/pokeapi"
	"github.com/ihaliti03/pokedex/internal/pokecache"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

type Config struct{
	Commands	map[string]CliCommand
	NextURL 	string
	PreviousURL	string
	Cache       *pokecache.Cache
	Pokedex     map[string]pokeapi.Pokemon
}

