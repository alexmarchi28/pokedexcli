package main

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/alexmarchi28/pokedexcli/internal/pokecache"
)

func TestInspectPokemonRequiresCaughtPokemon(t *testing.T) {
	cfg := &config{
		Pokedex: make(map[string]Pokemon),
	}

	output := captureOutput(t, func() {
		err := inspectPokemon(cfg, "pidgey")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	expected := "you have not caught that pokemon\n"
	if output != expected {
		t.Fatalf("expected %q, got %q", expected, output)
	}
}

func TestInspectPokemonUsesCacheForCaughtPokemon(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Minute)
	cache.Add(pokemonURL+"pidgey", []byte(`{
		"name": "pidgey",
		"base_experience": 50,
		"height": 3,
		"weight": 18,
		"stats": [
			{
				"base_stat": 40,
				"stat": {
					"name": "hp"
				}
			},
			{
				"base_stat": 45,
				"stat": {
					"name": "attack"
				}
			}
		],
		"types": [
			{
				"type": {
					"name": "normal"
				}
			},
			{
				"type": {
					"name": "flying"
				}
			}
		]
	}`))

	cfg := &config{
		Cache: cache,
		Pokedex: map[string]Pokemon{
			"pidgey": {
				Name:   "pidgey",
				Height: 999,
				Weight: 999,
			},
		},
	}

	output := captureOutput(t, func() {
		err := inspectPokemon(cfg, "pidgey")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	expected := `Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
Types:
  - normal
  - flying
`
	if output != expected {
		t.Fatalf("expected %q, got %q", expected, output)
	}
}

func TestInspectPokemonUsesPokedexWhenCacheMisses(t *testing.T) {
	cfg := &config{
		Pokedex: map[string]Pokemon{
			"pidgey": {
				Name:   "pidgey",
				Height: 3,
				Weight: 18,
				Stats: []PokemonStat{
					{Name: "hp", Value: 40},
				},
				Types: []string{"normal"},
			},
		},
	}

	output := captureOutput(t, func() {
		err := inspectPokemon(cfg, "pidgey")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	expected := `Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
Types:
  - normal
`
	if output != expected {
		t.Fatalf("expected %q, got %q", expected, output)
	}
}

func captureOutput(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("couldn't create pipe: %v", err)
	}

	os.Stdout = writer
	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("couldn't close writer: %v", err)
	}

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("couldn't read output: %v", err)
	}

	return string(output)
}
