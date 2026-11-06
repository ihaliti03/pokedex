package main

import (
	"time"

	"github.com/ihaliti03/pokedex/internal/cli"
	"github.com/ihaliti03/pokedex/internal/commands"
	"github.com/ihaliti03/pokedex/internal/pokeapi"
	"github.com/ihaliti03/pokedex/internal/pokecache"
)

func main(){
	cfg := &cli.Config{
        Commands: commands.BuildRegistry(),
		Cache:    pokecache.NewCache(5 * time.Minute),
		Pokedex:  make(map[string]pokeapi.Pokemon),
    }
	startRepl(cfg)
}