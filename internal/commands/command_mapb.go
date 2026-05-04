package commands

import (
	"errors"
	"fmt"

	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

func ShowPreviousMapLocations(cfg *Config, _ ...string) error {
	if cfg.Previous == "" {
		return errors.New("you're on the first page")
	}

	page, err := utils.GetLocationAreaPage(cfg.Previous, cfg.Cache)
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
