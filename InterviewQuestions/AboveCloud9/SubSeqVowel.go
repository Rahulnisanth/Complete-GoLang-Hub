package main

import "strings"

// Find all subsequences that start with a vowel
func FindSubSeqVowels(input string) []string {
	// Created string to check the rune presence
	vowels := "aeiouAEIOU"
	// Slice to store sub-sequences
	result := []string{}

	for i := 0; i < len(input); i++ {
		// Check the rune presence using vowels string
		if strings.ContainsRune(vowels, rune(input[i])) {
			for j := i + 1; j <= len(input); j++ {
				// Appended the sub-sequences using string slicing
				result = append(result, input[i:j])
			}
		}
	}
	return result
}
