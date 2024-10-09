// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.
package main

import "strconv"

// calculate() returns the sum of the two parameters
func calculateAnswer(value1 string, value2 string) float64 {
	// Convert the first string to a float64
	numFloat1, _ := strconv.ParseFloat(value1, 64)

	// Convert the second string to a float64
	numFloat2, _ := strconv.ParseFloat(value2, 64)

	// Calculate and return the result
	result := (numFloat1 + numFloat2)

	return result
}
