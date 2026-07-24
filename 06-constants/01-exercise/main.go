/*
Exercise 1: Month Enumeration using Go 'iota'

Demonstrate defining sequential local constant values using the 'iota' identifier.

Requirements:
- Declare a block of local constants representing months of the year.
- Use 'iota' to auto-increment values sequentially starting from 1 (January = 1).
- Print the numeric values corresponding to the first three months of the year to the console.
*/

package main

import "fmt"

func main() {
	// Declare a block of local constants starting at 1.
	// 'iota' resets to 0 at the start of a const block, so 'iota + 1' offsets the starting value to 1.
	const (
		january = iota + 1
		february
		march
	)

	// Print the numeric representations of the enumerated months.
	fmt.Println("=== Month Enumeration Report ===")
	fmt.Printf("January:  Month #%d\n", january)
	fmt.Printf("February: Month #%d\n", february)
	fmt.Printf("March:    Month #%d\n", march)
}
