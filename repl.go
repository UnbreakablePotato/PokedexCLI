package main

import (
	"strings"
)

func cleanInput(text string) []string {
	res := []string{}

	text = strings.ToLower(text)

	text = strings.TrimSpace(text)

	cleaned := strings.ReplaceAll(text, ",", "")

	inter := strings.Split(cleaned, " ")

	res = append(res, inter...)

	return res
}
