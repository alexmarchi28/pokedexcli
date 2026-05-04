package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/alexmarchi28/pokedexcli/internal/pokecache"
)

const LocationAreaURL = "https://pokeapi.co/api/v2/location-area/"

type LocationAreaPage struct {
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

type LocationAreaDetails struct {
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

func GetLocationAreaPage(url string, cache *pokecache.Cache) (LocationAreaPage, error) {
	if cache != nil {
		body, ok := cache.Get(url)
		if ok {
			return ParseLocationAreaPage(body)
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaPage{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaPage{}, err
	}

	if res.StatusCode > 299 {
		return LocationAreaPage{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	page, err := ParseLocationAreaPage(body)
	if err != nil {
		return LocationAreaPage{}, err
	}

	if cache != nil {
		cache.Add(url, body)
	}

	return page, nil
}

func ParseLocationAreaPage(body []byte) (LocationAreaPage, error) {
	var locRes locationAreaResponse
	if err := json.Unmarshal(body, &locRes); err != nil {
		return LocationAreaPage{}, err
	}

	names := make([]string, 0, len(locRes.Results))
	for _, location := range locRes.Results {
		names = append(names, location.Name)
	}

	return LocationAreaPage{
		Next:     locRes.Next,
		Previous: locRes.Previous,
		Names:    names,
	}, nil
}

func GetLocationAreaDetails(name string, cache *pokecache.Cache) (LocationAreaDetails, error) {
	locationURL := LocationAreaURL + url.PathEscape(name)

	if cache != nil {
		body, ok := cache.Get(locationURL)
		if ok {
			return ParseLocationAreaDetails(body)
		}
	}

	res, err := http.Get(locationURL)
	if err != nil {
		return LocationAreaDetails{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	if res.StatusCode == http.StatusNotFound {
		return LocationAreaDetails{}, fmt.Errorf("location area %q not found", name)
	}

	if res.StatusCode > 299 {
		return LocationAreaDetails{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	details, err := ParseLocationAreaDetails(body)
	if err != nil {
		return LocationAreaDetails{}, err
	}

	if cache != nil {
		cache.Add(locationURL, body)
	}

	return details, nil
}

func ParseLocationAreaDetails(body []byte) (LocationAreaDetails, error) {
	var locRes locationAreaDetailsResponse
	if err := json.Unmarshal(body, &locRes); err != nil {
		return LocationAreaDetails{}, err
	}

	names := make([]string, 0, len(locRes.PokemonEncounters))
	for _, encounter := range locRes.PokemonEncounters {
		if encounter.Pokemon.Name == "" {
			continue
		}
		names = append(names, encounter.Pokemon.Name)
	}

	return LocationAreaDetails{
		Name:         locRes.Name,
		PokemonNames: names,
	}, nil
}
