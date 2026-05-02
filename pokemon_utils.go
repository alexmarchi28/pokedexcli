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
}

type pokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
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

func parsePokemon(body []byte) (Pokemon, error) {
	var pokemonRes pokemonResponse
	if err := json.Unmarshal(body, &pokemonRes); err != nil {
		return Pokemon{}, err
	}

	return Pokemon{
		Name:           pokemonRes.Name,
		BaseExperience: pokemonRes.BaseExperience,
	}, nil
}
