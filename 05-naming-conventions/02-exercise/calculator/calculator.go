package calculator

import "math"

// Pi is a public (exported) constant representing the mathematical constant π.
const Pi float64 = 3.1416

// numberOfCalculationsPerformed is a private (unexported) counter tracking invocation count.
var numberOfCalculationsPerformed int = 0

// CalculateCircleArea increments the internal execution counter and calculates a circle's area.
func CalculateCircleArea(radius float64) float64 {
	numberOfCalculationsPerformed++
	return Pi * math.Pow(radius, 2)
}

// GetTotalCalculations is a public getter allowing external packages to read the internal counter.
func GetTotalCalculations() int {
	return numberOfCalculationsPerformed
}
