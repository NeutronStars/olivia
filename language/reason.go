package language

import (
	"strings"
)

// SearchReason returns the reason found in the given sentence for the reminders,
// here is an example: "Remind me that I need to **call mom** tomorrow".
func SearchReason(sentence string) string {
	var response []string

	// Split the given sentence into words
	words := strings.Split(sentence, " ")

	// Initialize the appeared boolean for the keywords
	appeared := false
	// Iterate through the words
	for _, word := range words {
		// Add the word to the response array if the keyword already appeared
		if appeared {
			response = append(response, word)
		}

		// If the keyword didn't appeared and one of the keywords match set the appeared condition
		// to true
		if !appeared && (LevenshteinDistance(word, "that") <= 2 ||
			LevenshteinDistance(word, "to") < 2) {
			appeared = true
		}
	}

	// Join the words and return the sentence
	return strings.Join(response, " ")
}
