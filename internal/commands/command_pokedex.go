package commands

import (
	"fmt"
	"sort"
)

func ShowPokedex(cfg *Config, _ ...string) error {
	fmt.Println("Your Pokedex:")

	names := make([]string, 0, len(cfg.Pokedex))
	for name := range cfg.Pokedex {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
