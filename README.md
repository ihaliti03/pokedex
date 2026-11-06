# Pokedex CLI

A command-line Pokedex built in Go. Explore location areas, discover Pokémon, catch them, and inspect your collection — all from an interactive REPL.

Built as part of a [Boot.dev](https://www.boot.dev) course.

## Features

- **Interactive REPL** — type commands, get responses, repeat
- **Location pagination** — page forward and backward through 1,000+ location areas
- **Explore areas** — see which Pokémon live in a given location area
- **Catch Pokémon** — random chance based on the Pokémon's base experience
- **Inspect caught Pokémon** — view height, weight, stats, and types
- **Pokedex** — list everything you've caught
- **Response caching** — API responses cached in memory with a TTL (5 minutes default)

## Requirements

- Go 1.21+ (uses generics, `slices`/`maps` from stdlib, auto-seeded `math/rand`)
- Internet connection (talks to [PokeAPI](https://pokeapi.co))

## Installation

```bash
git clone https://github.com/your-username/pokedex.git
cd pokedex
go build -o pokedex .