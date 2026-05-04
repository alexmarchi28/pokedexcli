package commands

import (
	"github.com/alexmarchi28/pokedexcli/internal/pokecache"
	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Pokedex  map[string]utils.Pokemon
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Show usage",
			Callback:    ShowHelp,
		},
		"map": {
			Name:        "map",
			Description: "Show 20 map locations",
			Callback:    ShowMapLocations,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show previous 20 map locations",
			Callback:    ShowPreviousMapLocations,
		},
		"explore": {
			Name:        "explore",
			Description: "Show pokemon in a location area",
			Callback:    ExploreLocationArea,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a pokemon",
			Callback:    CatchPokemon,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Show details for a caught pokemon",
			Callback:    InspectPokemon,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show caught pokemon",
			Callback:    ShowPokedex,
		},
	}
}
