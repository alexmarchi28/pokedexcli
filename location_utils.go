package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/alexmarchi28/pokedexcli/internal/pokecache"
)

const locationAreaURL = "https://pokeapi.co/api/v2/location-area/"

type locationAreaPage struct {
	Next     string
	Previous string
	Names    []string
}

type locationAreaResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaDetails struct {
	Name         string
	PokemonNames []string
}

type locationAreaDetailsResponse struct {
	Name              string                     `json:"name"`
	PokemonEncounters []locationPokemonEncounter `json:"pokemon_encounters"`
}

type locationPokemonEncounter struct {
	Pokemon locationPokemon `json:"pokemon"`
}

type locationPokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getLocationAreaPage(url string, cache *pokecache.Cache) (locationAreaPage, error) {
	if cache != nil {
		body, ok := cache.Get(url)
		if ok {
			return parseLocationAreaPage(body)
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreaPage{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaPage{}, err
	}

	if res.StatusCode > 299 {
		return locationAreaPage{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	page, err := parseLocationAreaPage(body)
	if err != nil {
		return locationAreaPage{}, err
	}

	if cache != nil {
		cache.Add(url, body)
	}

	return page, nil
}

func parseLocationAreaPage(body []byte) (locationAreaPage, error) {
	var locRes locationAreaResponse
	if err := json.Unmarshal(body, &locRes); err != nil {
		return locationAreaPage{}, err
	}

	names := make([]string, 0, len(locRes.Results))
	for _, location := range locRes.Results {
		names = append(names, location.Name)
	}

	return locationAreaPage{
		Next:     locRes.Next,
		Previous: locRes.Previous,
		Names:    names,
	}, nil
}

func getLocationAreaDetails(name string, cache *pokecache.Cache) (locationAreaDetails, error) {
	locationURL := locationAreaURL + url.PathEscape(name)

	if cache != nil {
		body, ok := cache.Get(locationURL)
		if ok {
			return parseLocationAreaDetails(body)
		}
	}

	res, err := http.Get(locationURL)
	if err != nil {
		return locationAreaDetails{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaDetails{}, err
	}

	if res.StatusCode == http.StatusNotFound {
		return locationAreaDetails{}, fmt.Errorf("location area %q not found", name)
	}

	if res.StatusCode > 299 {
		return locationAreaDetails{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	details, err := parseLocationAreaDetails(body)
	if err != nil {
		return locationAreaDetails{}, err
	}

	if cache != nil {
		cache.Add(locationURL, body)
	}

	return details, nil
}

func parseLocationAreaDetails(body []byte) (locationAreaDetails, error) {
	var locRes locationAreaDetailsResponse
	if err := json.Unmarshal(body, &locRes); err != nil {
		return locationAreaDetails{}, err
	}

	names := make([]string, 0, len(locRes.PokemonEncounters))
	for _, encounter := range locRes.PokemonEncounters {
		if encounter.Pokemon.Name == "" {
			continue
		}
		names = append(names, encounter.Pokemon.Name)
	}

	return locationAreaDetails{
		Name:         locRes.Name,
		PokemonNames: names,
	}, nil
}
