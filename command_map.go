package main

import (
	"encoding/json"
	"fmt"
)

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config, client *client, s string) error {
	baseUrl := ""

	if cfg.next != nil {
		baseUrl = *cfg.next
	} else {
		baseUrl = "https://pokeapi.co/api/v2/location-area?limit=20"
	}

	err := PrintAndPass(cfg, baseUrl, client)

	if err != nil {
		return err
	}

	return nil
}
func commandPreMap(cfg *config, client *client, s string) error {
	baseUrl := ""

	if cfg.previous != nil {
		baseUrl = *cfg.previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	err := PrintAndPass(cfg, baseUrl, client)

	if err != nil {
		return err
	}

	return nil
}

func PrintAndPass(cfg *config, baseUrl string, client *client) error {

	bytes, err := request(baseUrl, client)

	if err != nil {
		return err
	}
	locs, err := unMarshalOnLoc(bytes)

	if err != nil {
		return err
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	cfg.previous = locs.Previous
	cfg.next = locs.Next

	return nil
}

func unMarshalOnLoc(body []byte) (locationAreaResponse, error) {
	loc := locationAreaResponse{}

	err := json.Unmarshal(body, &loc)

	if err != nil {
		return locationAreaResponse{}, fmt.Errorf("unmarsh fail: %w", err)
	}

	return loc, nil
}
