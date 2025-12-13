package main

import (
	"encoding/json"
	"fmt"
)

func commandExplore(cfg *config, c *client, s string) error {
	if s == "" {
		return fmt.Errorf("missing location")
	}
	fmt.Printf("Exploring %s", s)
	fmt.Println("...")

	url := "https://pokeapi.co/api/v2/location-area/" + s

	data, err := request(url, c)
	if err != nil {
		return fmt.Errorf("failed request: %w", err)
	}
	areaData, err := unMarshalOnLocArea(data)
	if err != nil {
		return fmt.Errorf("failed unmarshal: %w", err)
	}

	if len(areaData.PokemonEncounters) == 0 {
		fmt.Println("No pokemon in this area")
		return nil
	}

	for _, encounter := range areaData.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
func unMarshalOnLocArea(body []byte) (locationArea, error) {
	loc := locationArea{}

	err := json.Unmarshal(body, &loc)

	if err != nil {
		return locationArea{}, fmt.Errorf("unmarsh fail: %w", err)
	}

	return loc, nil
}

type locationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
