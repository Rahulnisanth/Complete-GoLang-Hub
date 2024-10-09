// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"strings"
)

type cartItem struct {
	Name  string
	Price float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	// Slice declaration
	cart := make([]cartItem, 0, 20)
    // Your code goes here.
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}
    var item cartItem
	for decoder.More() {
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		cart = append(cart, item)
	}
	return cart
}
