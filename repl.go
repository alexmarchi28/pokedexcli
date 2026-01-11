package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	toLower := strings.ToLower(text)
	words := strings.Fields(toLower)
	return words
}

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

		firstWord := cleanedInput[0]

		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
