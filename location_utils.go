package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
