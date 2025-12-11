package main

import (
	"strings"
)

func cleanInput(text string) []string {
	newText := strings.Fields(strings.TrimSpace(strings.ToLower(text)))

	return newText
}
