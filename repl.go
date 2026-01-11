package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
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

		err := command.callback()
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
	callback    func() error
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
	}
}
