// Write your answer here, and then test your code.
package main

import "strconv"

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
    // Your code goes here.
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
		case "+":
			return addValues(value1, value2)
		case "-":
			return subtractValues(value1, value2)
		case "*":
			return multiplyValues(value1, value2)
		case "/":
			return divideValues(value1, value2)
		default:
			panic("Invalid operation")
	}
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
