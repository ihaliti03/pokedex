package commands

import (
	"fmt"
	"os"
	"github.com/ihaliti03/pokedex/internal/cli"
)

func CommandExit(_ *cli.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}