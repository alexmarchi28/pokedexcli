package main

import "testing"

func TestParsePokemon(t *testing.T) {
	body := []byte(`{
		"name": "pikachu",
		"base_experience": 112
	}`)

	actual, err := parsePokemon(body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actual.Name != "pikachu" {
		t.Fatalf("expected name pikachu, got %s", actual.Name)
	}

	if actual.BaseExperience != 112 {
		t.Fatalf("expected base experience 112, got %d", actual.BaseExperience)
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
			actual := catchChancePercentage(c.baseExperience)
			if actual != c.expected {
				t.Fatalf("expected %d, got %d", c.expected, actual)
			}
		})
	}
}
