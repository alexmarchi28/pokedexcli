package tests

import (
	"testing"

	"github.com/alexmarchi28/pokedexcli/internal/commands"
	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

func TestShowPokedexListsCaughtPokemon(t *testing.T) {
	cfg := &commands.Config{
		Pokedex: map[string]utils.Pokemon{
			"pidgey":  {Name: "pidgey"},
			"pikachu": {Name: "pikachu"},
		},
	}

	output := captureOutput(t, func() {
		err := commands.ShowPokedex(cfg)
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
	cfg := &commands.Config{}

	output := captureOutput(t, func() {
		err := commands.ShowPokedex(cfg)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	expected := "Your Pokedex:\n"
	if output != expected {
		t.Fatalf("expected %q, got %q", expected, output)
	}
}
