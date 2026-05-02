package main

import (
	"errors"
	"fmt"
	"strings"
)

func inspectPokemon(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := strings.Join(args, "-")
	if cfg.Pokedex == nil {
		cfg.Pokedex = make(map[string]Pokemon)
	}

	caughtPokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	pokemon := caughtPokemon
	cachedPokemon, ok := getCachedPokemon(pokemonName, cfg.Cache)
	if ok {
		pokemon = cachedPokemon
	}

	printPokemonDetails(pokemon)
	return nil
}

func printPokemonDetails(pokemon Pokemon) {
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
