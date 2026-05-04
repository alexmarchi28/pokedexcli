package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

func InspectPokemon(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := strings.Join(args, "-")
	if cfg.Pokedex == nil {
		cfg.Pokedex = make(map[string]utils.Pokemon)
	}

	caughtPokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	pokemon := caughtPokemon
	cachedPokemon, ok := utils.GetCachedPokemon(pokemonName, cfg.Cache)
	if ok {
		pokemon = cachedPokemon
	}

	printPokemonDetails(pokemon)
	return nil
}

func printPokemonDetails(pokemon utils.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Name, stat.Value)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType)
	}
}
