package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/alexmarchi28/pokedexcli/internal/pokecache"
)

const pokemonURL = "https://pokeapi.co/api/v2/pokemon/"

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []PokemonStat
	Types          []string
}

type PokemonStat struct {
	Name  string
	Value int
}

type pokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func getPokemon(name string, cache *pokecache.Cache) (Pokemon, error) {
	pokemonEndpointURL := pokemonURL + url.PathEscape(name)

	if cache != nil {
		body, ok := cache.Get(pokemonEndpointURL)
		if ok {
			return parsePokemon(body)
		}
	}

	res, err := http.Get(pokemonEndpointURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if res.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("pokemon %q not found", name)
	}

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	pokemon, err := parsePokemon(body)
	if err != nil {
		return Pokemon{}, err
	}

	if cache != nil {
		cache.Add(pokemonEndpointURL, body)
	}

	return pokemon, nil
}

func getCachedPokemon(name string, cache *pokecache.Cache) (Pokemon, bool) {
	if cache == nil {
		return Pokemon{}, false
	}

	pokemonEndpointURL := pokemonURL + url.PathEscape(name)
	body, ok := cache.Get(pokemonEndpointURL)
	if !ok {
		return Pokemon{}, false
	}

	pokemon, err := parsePokemon(body)
	if err != nil {
		return Pokemon{}, false
	}

	return pokemon, true
}

func parsePokemon(body []byte) (Pokemon, error) {
	var pokemonRes pokemonResponse
	if err := json.Unmarshal(body, &pokemonRes); err != nil {
		return Pokemon{}, err
	}

	stats := make([]PokemonStat, 0, len(pokemonRes.Stats))
	for _, stat := range pokemonRes.Stats {
		if stat.Stat.Name == "" {
			continue
		}
		stats = append(stats, PokemonStat{
			Name:  stat.Stat.Name,
			Value: stat.BaseStat,
		})
	}

	types := make([]string, 0, len(pokemonRes.Types))
	for _, pokemonType := range pokemonRes.Types {
		if pokemonType.Type.Name == "" {
			continue
		}
		types = append(types, pokemonType.Type.Name)
	}

	return Pokemon{
		Name:           pokemonRes.Name,
		BaseExperience: pokemonRes.BaseExperience,
		Height:         pokemonRes.Height,
		Weight:         pokemonRes.Weight,
		Stats:          stats,
		Types:          types,
	}, nil
}
