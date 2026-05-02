package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexmarchi28/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		Next:  locationAreaURL,
		Cache: pokecache.NewCache(5 * time.Minute),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		rawInput := scanner.Text()
		cleanedInput := cleanInput(rawInput)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		command, exists := getCommands()[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	toLower := strings.ToLower(text)
	words := strings.Fields(toLower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show usage",
			callback:    showHelp,
		},
		"map": {
			name:        "map",
			description: "Show 20 map locations",
			callback:    showMapLocations,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 map locations",
			callback:    showPreviousMapLocations,
		},
	}
}
