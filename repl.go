package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ihaliti03/pokedex/internal/cli"
)

func startRepl(cfg *cli.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan(){
			break
		}
		// v2
		cleaned := cleanInput(scanner.Text())

		if len(cleaned) == 0 {
			continue
		}

		cmd, ok := cfg.Commands[cleaned[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.Callback(cfg, cleaned[1:]); err != nil {
			fmt.Println("error:", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func cleanInput(text string) []string{
	// instead of Split on " " this is great
	result := strings.Fields(strings.ToLower(text))
	// if we used raw split function we would need to add additional logic to this function
	// to fix the issue of one more slice element of value "" so it doesnt return last space as a slice element
	return result
}
