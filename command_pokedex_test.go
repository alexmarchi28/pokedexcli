package main

import "testing"

func TestShowPokedexListsCaughtPokemon(t *testing.T) {
	cfg := &config{
		Pokedex: map[string]Pokemon{
			"pidgey":  {Name: "pidgey"},
			"pikachu": {Name: "pikachu"},
		},
	}

	output := captureOutput(t, func() {
		err := showPokedex(cfg)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	expected := `Your Pokedex:
 - pidgey
 - pikachu
`
	if output != expected {
		t.Fatalf("expected %q, got %q", expected, output)
	}
}

func TestShowPokedexHandlesNoCaughtPokemon(t *testing.T) {
	cfg := &config{}

	output := captureOutput(t, func() {
		err := showPokedex(cfg)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	expected := "Your Pokedex:\n"
	if output != expected {
		t.Fatalf("expected %q, got %q", expected, output)
	}
}
