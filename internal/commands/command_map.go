package commands

import (
	"fmt"

	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

func ShowMapLocations(cfg *Config, _ ...string) error {
	if cfg.Next == "" {
		cfg.Next = utils.LocationAreaURL
	}

	page, err := utils.GetLocationAreaPage(cfg.Next, cfg.Cache)
	if err != nil {
		return err
	}

	for _, name := range page.Names {
		fmt.Println(name)
	}

	cfg.Next = page.Next
	cfg.Previous = page.Previous

	return nil
}
