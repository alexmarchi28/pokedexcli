package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

func ExploreLocationArea(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location area")
	}

	locationAreaName := strings.Join(args, "-")

	details, err := utils.GetLocationAreaDetails(locationAreaName, cfg.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", details.Name)
	for _, name := range details.PokemonNames {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
