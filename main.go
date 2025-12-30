package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	ctx := Context{
		ResultsPerPage: 20,
		Client:         pokeapi.NewClient(time.Hour),
	}
	replLoop(&ctx)
}
