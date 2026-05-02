package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func catchPokemon(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := strings.Join(args, "-")
	pokemon, err := getPokemon(pokemonName, cfg.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if !pokemonWasCaught(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	if cfg.Pokedex == nil {
		cfg.Pokedex = make(map[string]Pokemon)
	}
	cfg.Pokedex[pokemon.Name] = pokemon

	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}

func pokemonWasCaught(baseExperience int) bool {
	return rand.Intn(100) < catchChancePercentage(baseExperience)
}

func catchChancePercentage(baseExperience int) int {
	const (
		maxChance                 = 80
		minChance                 = 30
		baseExperienceForMinCatch = 600
	)

	if baseExperience <= 0 {
		return maxChance
	}

	chance := maxChance - (baseExperience * (maxChance - minChance) / baseExperienceForMinCatch)
	if chance < minChance {
		return minChance
	}
	if chance > maxChance {
		return maxChance
	}
	return chance
}
