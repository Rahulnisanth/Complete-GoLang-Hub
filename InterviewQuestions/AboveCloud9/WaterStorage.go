package main

import (
	"fmt"
)

// Function to calculate the water storage
func waterStorage(heights []int) {
	n := len(heights)
	if n < 3 {
		fmt.Println("Impossible to store water")
	} else {
		result := 0
		leftMax := make([]int, n)
		rightMax := make([]int, n)

		// Fill the left max array for every indices
		leftMax[0] = heights[0]
		for i := 1; i < n; i++ {
			leftMax[i] = max(leftMax[i-1], heights[i])
		}

		// Fill the right max array for every indices
		rightMax[n-1] = heights[n-1]
		for i := n - 2; i >= 0; i-- {
			rightMax[i] = max(rightMax[i+1], heights[i])
		}

		// Calculate the volume of water stored in between each blocks
		for i := 0; i < n; i++ {
			waterLevel := min(leftMax[i], rightMax[i])
			if waterLevel > heights[i] {
				result += waterLevel - heights[i]
			}
		}

		fmt.Print(result)
	}
}
