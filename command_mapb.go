package main

import (
	"errors"
	"fmt"
)

func showPreviousMapLocations(cfg *config) error {
	if cfg.Previous == "" {
		return errors.New("you're on the first page")
	}

	page, err := getLocationAreaPage(cfg.Previous)
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
