package commands

import "github.com/ihaliti03/pokedex/internal/cli"

func BuildRegistry() map[string]cli.CliCommand {
    return map[string]cli.CliCommand{
        "exit": {
			Name: "exit", 
			Description: "Exit the Pokedex", 
			Callback: CommandExit,
		},
        "help": {
			Name: "help", 
			Description: "Displays a help message", 
			Callback: CommandHelp,
		},
        "map":  {
			Name: "map",  
			Description: "Get next location areas", 
			Callback: CommandMap,
		},
        "mapb": {
			Name: "mapb", 
			Description: "Get previous location areas", 
			Callback: CommandMapB,
		},
		"explore": {
            Name:        "explore",
            Description: "Explore a location area for Pokémon",
            Callback:    CommandExplore,
        },
		"catch": {
		    Name:        "catch",
		    Description: "Catch a Pokémon by name",
		    Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
		    Description: "Display details of catched pokemon",
		    Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List caught Pokemon",
			Callback:    CommandPokedex,
		},
    }
}