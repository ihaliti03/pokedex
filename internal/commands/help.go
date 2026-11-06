package commands

import (
	"fmt"
	"github.com/ihaliti03/pokedex/internal/cli"
)

func CommandHelp(cfg *cli.Config,args []string) error {
	printUsage(cfg)
	return nil
}

func printUsage(cfg *cli.Config) {
	fmt.Println("Usage:")
	fmt.Println("")

	for _, cmd := range cfg.Commands {
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
    }
}