package main

import (
	"pokedexcli/internal/pokeapi"
)

type Context struct {
	Next           string
	Previous       string
	ResultsPerPage int
	Client         pokeapi.Client
}
