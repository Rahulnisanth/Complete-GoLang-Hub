package main

import (
	"fmt"
)

func main() {
	// Prompt the length
	var n int
	fmt.Scanln(&n)

	// Prompt the block heights
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&l[i])
	}

	// Core logic
	if n < 3 {
		fmt.Println("Impossible to store water")
	} else {
		result := 0
		leftMax := make([]int, n)
		rightMax := make([]int, n)

		// Fill the left max array for every indices
		leftMax[0] = l[0]
		for i := 1; i < n; i++ {
			leftMax[i] = max(leftMax[i-1], l[i])
		}

		// Fill the right max array for every indices
		rightMax[n-1] = l[n-1]
		for i := n - 2; i >= 0; i-- {
			rightMax[i] = max(rightMax[i+1], l[i])
		}

		// Calculate the volume of water stored in between each blocks
		for i := 0; i < n; i++ {
			waterLevel := min(leftMax[i], rightMax[i])
			if waterLevel > l[i] {
				result += waterLevel - l[i]
			}
		}

		fmt.Print(result)
	}
}
