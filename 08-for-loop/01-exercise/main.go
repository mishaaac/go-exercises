/*
Exercise 1: Sum of Elements at Even Indexes

Write a program that iterates over a slice of integers and calculates
the sum of elements located strictly at even index positions (0, 2, 4, ...).

Requirements:
- Define a slice populated with integer values.
- Iterate through the slice keeping track of both index and element value.
- Filter elements based on whether their index is even using the modulo operator.
- Output the accumulated total in a clear, formatted message.
*/

package main

import "fmt"

func main() {
	// Initialize a slice of integers representing the input data set.
	numbers := []int{10, 15, 20, 25, 30}

	// Track the cumulative sum of values found at even index locations.
	evenIndexSum := 0

	// Iterate through the slice while keeping track of both
	// the position (index) and its corresponding value.
	for index, value := range numbers {

		// Check if the current position is an even index.
		// Zero is considered an even index, so index 0 is included.
		if index%2 == 0 {

			// Accumulate only values that belong to even index positions.
			evenIndexSum += value
		}
	}

	// Print the result with clear and concise phrasing.
	fmt.Printf("Sum of values at even indexes: %d\n", evenIndexSum)
}
