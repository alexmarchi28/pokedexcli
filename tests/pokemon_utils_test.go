package tests

import (
	"testing"

	"github.com/alexmarchi28/pokedexcli/internal/commands"
	"github.com/alexmarchi28/pokedexcli/internal/utils"
)

func TestParsePokemon(t *testing.T) {
	body := []byte(`{
		"name": "pikachu",
		"base_experience": 112,
		"height": 4,
		"weight": 60,
		"stats": [
			{
				"base_stat": 35,
				"stat": {
					"name": "hp"
				}
			},
			{
				"base_stat": 55,
				"stat": {
					"name": "attack"
				}
			}
		],
		"types": [
			{
				"type": {
					"name": "electric"
				}
			}
		]
	}`)

	actual, err := utils.ParsePokemon(body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actual.Name != "pikachu" {
		t.Fatalf("expected name pikachu, got %s", actual.Name)
	}

	if actual.BaseExperience != 112 {
		t.Fatalf("expected base experience 112, got %d", actual.BaseExperience)
	}

	if actual.Height != 4 {
		t.Fatalf("expected height 4, got %d", actual.Height)
	}

	if actual.Weight != 60 {
		t.Fatalf("expected weight 60, got %d", actual.Weight)
	}

	if len(actual.Stats) != 2 {
		t.Fatalf("expected 2 stats, got %d", len(actual.Stats))
	}

	if actual.Stats[0].Name != "hp" || actual.Stats[0].Value != 35 {
		t.Fatalf("expected hp stat 35, got %s %d", actual.Stats[0].Name, actual.Stats[0].Value)
	}

	if len(actual.Types) != 1 {
		t.Fatalf("expected 1 type, got %d", len(actual.Types))
	}

	if actual.Types[0] != "electric" {
		t.Fatalf("expected electric type, got %s", actual.Types[0])
	}
}

func TestCatchChancePercentage(t *testing.T) {
	cases := []struct {
		name           string
		baseExperience int
		expected       int
	}{
		{
			name:           "lowest base experience",
			baseExperience: 0,
			expected:       80,
		},
		{
			name:           "middle base experience",
			baseExperience: 300,
			expected:       55,
		},
		{
			name:           "high base experience",
			baseExperience: 600,
			expected:       30,
		},
		{
			name:           "above max scaling",
			baseExperience: 900,
			expected:       30,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := commands.CatchChancePercentage(c.baseExperience)
			if actual != c.expected {
				t.Fatalf("expected %d, got %d", c.expected, actual)
			}
		})
	}
}
