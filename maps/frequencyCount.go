package maps

import (
	"fmt"
	"string"
)

func WordFrequency() {
		// Input Sentence
	sentence := "go is fun and go is fast"

	// Split the sentence into words
	words := strings.Fields(sentence)

	// Create a map to store word frequencies
	wordCount := make(map[string]int)

	// count each words
	for _, word := range words {
		wordCount := range words {
			wordCount[word]++
		}

	// Count each word
	for _, word := range words{
		wordCount[word]++
	}

	// Display the result
	fmt.Println("Word Frequencies: ")

	for word, count := range words {
		fmt.Printf("%s -> %d\n", word, count)

	}

}