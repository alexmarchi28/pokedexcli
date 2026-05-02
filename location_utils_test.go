package main

import "testing"

func TestParseLocationAreaDetails(t *testing.T) {
	body := []byte(`{
		"name": "canalave-city-area",
		"pokemon_encounters": [
			{
				"pokemon": {
					"name": "tentacool",
					"url": "https://pokeapi.co/api/v2/pokemon/72/"
				}
			},
			{
				"pokemon": {
					"name": "magikarp",
					"url": "https://pokeapi.co/api/v2/pokemon/129/"
				}
			}
		]
	}`)

	actual, err := parseLocationAreaDetails(body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actual.Name != "canalave-city-area" {
		t.Fatalf("expected name canalave-city-area, got %s", actual.Name)
	}

	expectedPokemon := []string{"tentacool", "magikarp"}
	if len(actual.PokemonNames) != len(expectedPokemon) {
		t.Fatalf("expected %d pokemon, got %d", len(expectedPokemon), len(actual.PokemonNames))
	}

	for i, expected := range expectedPokemon {
		if actual.PokemonNames[i] != expected {
			t.Errorf("expected pokemon %s at index %d, got %s", expected, i, actual.PokemonNames[i])
		}
	}
}
