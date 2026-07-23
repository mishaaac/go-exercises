/*
Exercise 3: Console Temperature Converter (Celsius to Fahrenheit)

Write a Go program that converts Celsius temperatures to Fahrenheit using mathematical
formulas and formats the output into a clean tabular layout.

Formula:

	Fahrenheit = (Celsius * 9 / 5) + 32

Requirements:
- Define two Celsius temperature floating-point variables (e.g., 0.0 and 100.0).
- Calculate their Fahrenheit equivalents using standard arithmetic operators.
- Format floating-point numbers to exactly 2 decimal places (`%.2f`).
- Display formatted output structured into a table view.
*/

package main

import "fmt"

func main() {
	// Temperature variables declared using short declaration syntax (:="").
	celsius1 := 0.0
	celsius2 := 100.0

	// Celsius to Fahrenheit conversion math: F = (C * 9/5) + 32
	fahrenheit1 := (celsius1 * 9.0 / 5.0) + 32.0
	fahrenheit2 := (celsius2 * 9.0 / 5.0) + 32.0

	// Tabular console header
	fmt.Println("[ Celsius ] -> [ Fahrenheit ]")
	fmt.Println("-----------------------------")

	// Print formatted floating-point values rounded to 2 decimal places (%.2f)
	fmt.Printf("%-10.2f °C -> %.2f °F\n", celsius1, fahrenheit1)
	fmt.Printf("%-10.2f °C -> %.2f °F\n", celsius2, fahrenheit2)
}
