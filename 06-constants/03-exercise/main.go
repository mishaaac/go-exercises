/*
Exercise 3: Content Delivery Network (CDN) Region Dispatcher

Design the core identifier system for a CDN dispatcher using arithmetic expressions with 'iota'.

Requirements:
- Define a custom integer-based type 'ServerRegion'.
- Declare global constants representing server deployment regions:
  * US_EAST, US_WEST, EU_CENTRAL, AP_SOUTHEAST.
- Strict constraint: 'US_EAST' must start at numeric value 100, with subsequent
  regions auto-incrementing by steps of 10 (100, 110, 120, 130...) using a single
  arithmetic expression with 'iota' inside the constant block.
- Output formatted region identifier records demonstrating correct step scaling.
*/

package main

import "fmt"

// ServerRegion defines a custom integer-based type for routing and telemetry identifiers.
type ServerRegion int

// Package-level constants representing geographic CDN edge regions.
// 'iota' scales linear constant sequences using arithmetic expressions:
// Formula: Base Value (100) + (iota * Step Increment (10))
const (
	UsEast      ServerRegion = 100 + (iota * 10) // iota = 0 -> 100 + (0 * 10) = 100
	UsWest                                       // iota = 1 -> 100 + (1 * 10) = 110
	EuCentral                                    // iota = 2 -> 100 + (2 * 10) = 120
	ApSoutheast                                  // iota = 3 -> 100 + (3 * 10) = 130
)

// String implements the fmt.Stringer interface to provide human-readable region names.
func (r ServerRegion) String() string {
	switch r {
	case UsEast:
		return "US-East (N. Virginia)"
	case UsWest:
		return "US-West (Oregon)"
	case EuCentral:
		return "EU-Central (Frankfurt)"
	case ApSoutheast:
		return "AP-Southeast (Singapore)"
	default:
		return fmt.Sprintf("Unknown Region Code (%d)", int(r))
	}
}

func main() {
	fmt.Println("=== CDN Edge Dispatcher Region Registry ===")

	// Output region identifiers showing numerical values and Stringer implementation.
	fmt.Printf("Region: %-25s | Region ID: %d\n", UsEast, UsEast)
	fmt.Printf("Region: %-25s | Region ID: %d\n", UsWest, UsWest)
	fmt.Printf("Region: %-25s | Region ID: %d\n", EuCentral, EuCentral)
	fmt.Printf("Region: %-25s | Region ID: %d\n", ApSoutheast, ApSoutheast)
}
