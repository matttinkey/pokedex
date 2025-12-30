package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Context, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get last 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Find pokemon in a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get information of a pokemon from pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all caught pokemon",
			callback:    commandPokedex,
		},
	}
}

var Pokedex map[string]pokeapi.Pokemon = map[string]pokeapi.Pokemon{}

func commandExit(ctx *Context, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(ctx *Context, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(ctx *Context, arg string) error {
	if ctx.Next == "" {
		ctx.Next = pokeapi.GetLocationsUrl(ctx.ResultsPerPage, 0)
	}

	url := ctx.Next

	locations, err := ctx.Client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	ctx.Previous = url
	ctx.Next = locations.Next
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapBack(ctx *Context, arg string) error {
	if ctx.Previous == "" {
		fmt.Println("Reached beginning, no more locations to display")
		return nil
	}
	url := ctx.Previous
	locations, err := ctx.Client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	ctx.Previous = locations.Previous
	ctx.Next = url
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExplore(ctx *Context, arg string) error {
	info, err := ctx.Client.GetLocationInfo(arg)
	if err != nil {
		return err
	}

	for _, encounter := range info.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(ctx *Context, arg string) error {
	pokemon, err := ctx.Client.GetPokemon(arg)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	num := rand.IntN(320)
	if num < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	Pokedex[pokemon.Name] = pokemon
	return nil
}

func commandInspect(ctx *Context, arg string) error {
	pokemon, ok := Pokedex[arg]
	if !ok {
		fmt.Printf("%s was not found in the pokedex\n", pokemon.Name)
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}

	return nil
}

func commandPokedex(ctx *Context, arg string) error {
	if Pokedex == nil {
		fmt.Println("No Pokemon have been caught")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k := range Pokedex {
		fmt.Printf("- %v\n", k)
	}
	return nil
}

func printLocations(locations pokeapi.LocationAreas) {
	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}
}
