package main

import (
	"fmt"
	"math"
)

func main() {
	// Define a number
	number := 123.456789

	// Round the number to the nearest integer
	rounded := math.Round(number * 100)

	// Print the rounded number with 2 decimal places
	fmt.Printf("%.3f\n", rounded) // Output: 12346.00

	// Divide the rounded number by 100 to get the original value rounded to 2 decimal places
	originalValue := rounded / 100
	fmt.Printf("%.3f\n", originalValue) // Output: 123.46
}
