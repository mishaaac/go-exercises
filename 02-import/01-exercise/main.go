/*
Exercise 1: Square Root Calculation with standard library math

Write a Go program that imports the standard library "math" package
to calculate and display the square root of the number 144.

Requirements:
- Import "fmt" and "math" standard library packages.
- Calculate square root using `math.Sqrt()`.
- Pass a `float64` argument to `math.Sqrt()` (it requires float64 inputs).
- Print formatted output displaying the input and calculated square root.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Sqrt requires a float64 parameter.
	number := 144.0

	// Calculate square root using standard library math package.
	result := math.Sqrt(number)

	// Display output using format specifiers (%.0f for clean whole number display).
	fmt.Printf("The square root of %.0f is %.0f\n", number, result)
}
