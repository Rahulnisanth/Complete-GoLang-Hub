package main

import (
	"fmt"
)

func main() {
	var output string
	j := 12.47
	if j < 12.0 { 
		output = "less" 
	} else if j > 13 {
		output = "more" 
	} else {
		output = "near" 
	} 
	fmt.Println(output) 
}
