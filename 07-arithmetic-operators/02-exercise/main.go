/*
Exercise 2: Fahrenheit to Celsius Temperature Converter

Write a program that converts temperatures from Fahrenheit to Celsius
using the physical conversion formula: C = (F - 32) * (5 / 9).

Requirements:
- Define a function 'fahrenheitToCelsius(fahrenheit float64) float64'.
- Avoid integer division pitfalls by using explicit floating-point literals (5.0 / 9.0).
- Format and display the output value with floating-point precision verbs (%.2f).
*/

package main

import "fmt"

// fahrenheitToCelsius converts a temperature reading from Fahrenheit to Celsius.
// Uses explicit float64 literals (5.0 / 9.0) to prevent integer truncation errors.
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * (5.0 / 9.0)
}

func main() {
	// Sample temperature reading in Fahrenheit.
	fahrenheitTemp := 100.0

	// Calculate Celsius equivalent.
	celsiusTemp := fahrenheitToCelsius(fahrenheitTemp)

	// Output formatted conversion report.
	fmt.Println("=== Temperature Conversion Report ===")
	fmt.Printf("Fahrenheit Input: %.2f°F\n", fahrenheitTemp)
	fmt.Printf("Celsius Output:   %.2f°C\n", celsiusTemp)
}
