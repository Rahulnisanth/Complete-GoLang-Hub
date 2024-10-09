package main

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

    // Create a map object
	result := make(map[string]float64)
	
    // Your code goes here
    for _, val := range items {
        result[val] = float64(1) / float64(len(items)) * 100
    }
    return result
}
