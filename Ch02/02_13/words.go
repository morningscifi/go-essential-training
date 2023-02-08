package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	for k, v := range wordCounter(strings.ToLower(text)) {
		fmt.Printf("%s %d\n", k, v)
	}
	fmt.Println(text)
}

func wordCounter(text string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range strings.Fields(text) {
		wordMap[word]++
	}

	return wordMap
}
