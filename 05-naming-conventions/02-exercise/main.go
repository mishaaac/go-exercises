/*
Exercise 2: Custom Package & Package Visibility

Design a package called 'calculator' that contains:
- A public constant for the value of Pi (3.1416).
- A private internal variable to store the total number of calculations performed (starting at 0).
- A public function 'CalculateCircleArea' that increments the internal counter and returns the area.

Requirements:
- Export public constant 'Pi' using PascalCase (`Pi = 3.1416`).
- Keep 'numberOfCalculationsPerformed' unexported using lowerCamelCase to ensure private package scope.
- Increment internal state counter on every invocation of 'CalculateCircleArea'.
- Import and consume package methods inside main.go for radius values 5 and 15.
*/

package main

import (
	"05-naming-conventions/02-exercise/calculator"
	"fmt"
)

func main() {
	radius1 := 5.0
	radius2 := 15.0

	area1 := calculator.CalculateCircleArea(radius1)
	area2 := calculator.CalculateCircleArea(radius2)

	fmt.Println("=== 📐 Circle Area Calculator ===")
	fmt.Printf("Pi constant value: %.4f\n\n", calculator.Pi)

	fmt.Printf("Area of circle with radius %.0f: %.2f\n", radius1, area1)
	fmt.Printf("Area of circle with radius %.0f: %.2f\n\n", radius2, area2)

	fmt.Printf("Total calculations executed: %d\n", calculator.GetTotalCalculations())
}
