/*
Exercise 2: Element Frequency Counter

Write a program that searches a slice of strings and calculates how many times
a specific target string appears within the collection.

Requirements:
- Define a slice of strings containing repeating elements.
- Define a target string to search for.
- Iterate through the slice using `for range`.
- Safely discard the unused index variable using the blank identifier (`_`).
- Maintain a tally counter that increments on exact matches.
- Display the target string and its total count in a professional format.
*/

package main

import "fmt"

func main() {
	// Initialize the data slice containing names with duplicate values.
	names := []string{"Alice", "Bob", "Alice", "Charlie", "Alice", "Bob"}

	// Define the search term to count.
	target := "Alice"

	// Track the total number of matches found during iteration.
	matchCount := 0

	// Traversal using range. The blank identifier '_' explicitly discards
	// the index since position data is irrelevant for frequency counting.
	for _, name := range names {

		// Compare current element against target for an exact match.
		if name == target {

			// Increment the match counter when a match occurs.
			matchCount++
		}
	}

	// Output the results with contextual information.
	fmt.Printf("Occurrences of %q: %d\n", target, matchCount)
}
