package main

import "fmt"

// Function to check if a string is a palindrome
func checkPalindrome(word string, start int, end int) bool {
	for start < end {
		if word[start] != word[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// Function to reverse a string
func reverseString(input string) string {
	i := 0
	j := len(input) - 1
	word := []rune(input)
	for i < j {
		word[i], word[j] = word[j], word[i]
		i++
		j--
	}
	return string(word)
}

// Function to make a string palindrome
func makePalindrome(s string) {
	if checkPalindrome(s, 0, len(s)-1) {
		fmt.Println(s)
		return
	}
	// Core logic
	N := len(s)
	revString := reverseString(s)
	for i := 0; i < N; i++ {
		if checkPalindrome(s+revString[N-i:], 0, len(s+revString[N-i:])-1) {
			fmt.Println(s + revString[N-i:])
			return
		}
	}
	fmt.Println("Impossible")
}

func main() {
	var s string
	fmt.Scanln(&s)
	makePalindrome(s)
}
