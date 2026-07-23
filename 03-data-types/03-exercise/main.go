/*
Exercise 3: Industrial Fuel Cylinder Filling Control System

Design a filling control system for an industrial fuel cylinder that:
- Receives a reading in liters (`float64`) and total capacity in cubic meters (`float64`).
- Converts capacity from cubic meters to liters.
- Calculates fill percentage accurately using floating-point math.
- Evaluates safety threshold conditions (e.g., overfill warning when > 95%).
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// Cylinder parameters
	litersRead := 1500.0         // Current fuel level in liters
	capacityInCubicMeters := 2.0 // Total cylinder capacity in cubic meters (m³)

	// Conversion formula: 1 cubic meter = 1000 liters
	capacityInLiters := capacityInCubicMeters * 1000.0

	// Calculate percentage fill (exact floating-point value)
	rawPercentage := (litersRead / capacityInLiters) * 100.0

	// Floor percentage to nearest whole integer for discrete readout display
	percentageFill := int(math.Floor(rawPercentage))

	// Safety check: trigger overfilled status if capacity exceeds 95%
	isOverfilled := percentageFill > 95

	fmt.Println("=== Industrial Fuel Control Panel ===")
	fmt.Printf("Current Fuel Level: %.1f L / %.1f L\n", litersRead, capacityInLiters)
	fmt.Printf("Fill Percentage   : %d%%\n", percentageFill)
	fmt.Printf("Overfilled Alert  : %t\n", isOverfilled)

	// Additional safety indicator
	if isOverfilled {
		fmt.Println("⚠️  WARNING: High fuel level! Stop filling immediately.")
	} else {
		fmt.Println("✅ Status: Fuel level within safe operating limit.")
	}
}
