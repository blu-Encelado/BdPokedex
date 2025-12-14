package main

import (
	pokeCache "BdPokedex/internal/cache"
	"fmt"
	"io"
	"net/http"
	"time"
)

type client struct {
	cache   *pokeCache.Cache
	pokedex map[string]pokemon
}

func NewClient() *client {
	return &client{
		cache:   pokeCache.NewCache(100 * time.Second),
		pokedex: make(map[string]pokemon),
	}
}

func request(url string, client *client) ([]byte, error) {

	if data, ok := client.cache.Get(url); ok {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fail to Get")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed: %w, status code: %d", err, res.StatusCode)
	}
	client.cache.Add(url, body)

	return body, nil
}

func addPokemonAtPokedex(data pokemon, c *client) string {
	_, ok := c.pokedex[data.Name]
	if ok {
		return fmt.Sprintf("you already got a %s", data.Name)
	}
	c.pokedex[data.Name] = data
	return fmt.Sprintf("%s was caught!", data.Name)
}
