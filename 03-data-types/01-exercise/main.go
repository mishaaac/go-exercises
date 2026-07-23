/*
Exercise 1: Type Conversion and Average Calculation

Write a Go program that calculates the average of three school grades.
Two grades are integers (`int`) and one is a decimal (`float64`).
The final result should be displayed as a decimal number formatted with
two decimal places.

Requirements:
- Define two integer variables (`int`) and one floating-point variable (`float64`).
- Perform explicit type casting (`float64(var)`) on integer values before arithmetic operations.
- Divide by floating-point literal `3.0` to avoid implicit type confusion.
- Output formatted result using `fmt.Printf("%.2f\n", average)`.
*/

package main

import "fmt"

func main() {
	// School grades using mixed types.
	grade1 := 8
	grade2 := 9
	grade3 := 7.5

	// Convert int values to float64 to perform precision floating-point division.
	// (8.0 + 9.0 + 7.5) / 3.0 = 24.5 / 3.0 = 8.1666...
	average := (float64(grade1) + float64(grade2) + grade3) / 3.0

	// Print result formatted to 2 decimal places (8.17).
	fmt.Printf("The final average of the grades is: %.2f\n", average)
}
