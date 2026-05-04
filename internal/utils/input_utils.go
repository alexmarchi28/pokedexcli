package utils

import "strings"

func CleanInput(text string) []string {
	toLower := strings.ToLower(text)
	words := strings.Fields(toLower)
	return words
}
