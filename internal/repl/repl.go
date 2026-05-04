package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/alexmarchi28/pokedexcli/internal/commands"
	"github.com/alexmarchi28/pokedexcli/internal/pokecache"
	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &commands.Config{
		Next:    utils.LocationAreaURL,
		Cache:   pokecache.NewCache(5 * time.Minute),
		Pokedex: make(map[string]utils.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		rawInput := scanner.Text()
		cleanedInput := utils.CleanInput(rawInput)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		command, exists := commands.GetCommands()[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.Callback(cfg, cleanedInput[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
