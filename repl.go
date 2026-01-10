package main

import "strings"

func cleanInput(text string) []string {
	toLower := strings.ToLower(text)
	words := strings.Fields(toLower)
	return words
}
