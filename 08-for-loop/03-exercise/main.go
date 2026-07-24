/*
Exercise 3: Target Pair Search (Two-Sum with Labeled Break)

Write a program that searches a slice of integers for the first pair of distinct
elements whose sum equals a specified target sum.
Use nested loops to evaluate pairs, and terminate all search processing as soon
as the first matching pair is discovered.

Requirements:
- Define a slice of integers (`numbers`) and a target integer (`targetSum = 15`).
- The outer loop selects the current element a[i], while the inner loop iterates over the remaining elements to its right (a[j], where j > i).
- Match condition: `numbers[i] + numbers[j] == targetSum`.
- Halt search execution immediately using a labeled break (`break Label`) or boolean flag logic.
- Output the discovered pair and their slice index positions, or print a fallback message if no match exists.
*/

package main

import "fmt"

func main() {
	// Source collection of integers.
	numbers := []int{10, 15, 3, 7, 8, 12}

	// Target value to match.
	targetSum := 15

	// Flag tracking search result.
	found := false

	fmt.Printf("Searching for target sum %d in slice: %v\n\n", targetSum, numbers)

	// Labeled loop permits breaking out of both nested loops simultaneously upon discovery.
SearchLoop:
	for i := 0; i < len(numbers); i++ {

		// Start inner index at i+1 to avoid comparing an element with itself
		// and to avoid testing duplicate pairs in reverse order.
		for j := i + 1; j < len(numbers); j++ {

			// Pair matching check.
			if numbers[i]+numbers[j] == targetSum {
				fmt.Printf("✅ Pair found: %d + %d = %d (Indices [%d] and [%d])\n",
					numbers[i], numbers[j], targetSum, i, j)

				found = true

				// Terminate BOTH loops immediately upon finding the first pair.
				break SearchLoop
			}
		}
	}

	// Report missing pair condition.
	if !found {
		fmt.Printf("❌ No pair found that sums to %d.\n", targetSum)
	}
}
