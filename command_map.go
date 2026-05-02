package main

import "fmt"

func showMapLocations(cfg *config) error {
	if cfg.Next == "" {
		cfg.Next = locationAreaURL
	}

	page, err := getLocationAreaPage(cfg.Next, cfg.Cache)
	if err != nil {
		return err
	}

	for _, name := range page.Names {
		fmt.Println(name)
	}

	cfg.Next = page.Next
	cfg.Previous = page.Previous

	return nil
}
