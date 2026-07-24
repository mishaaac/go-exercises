/*
Exercise 1: Rocket Launch Countdown (For as While)

Write a program that simulates a rocket launch countdown starting from 5 down to 1,
decreasing by 1 on each step. When the counter reaches 0, the program halts the countdown
and outputs "Liftoff!".

Requirements:
- Emulate a traditional `while` loop using Go's single-condition `for` loop (`for counter > 0`).
- Print the remaining seconds in "T-minus X seconds..." format on each iteration.
- Decrement the countdown variable on each iteration.
- Print "Liftoff!" once the loop completes when `counter == 0`.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Starting countdown value.
	counter := 5

	fmt.Println("=== Rocket Launch Sequence Initiated ===")

	// Single-condition loop acting as a 'while' loop.
	// Execution continues as long as counter remains strictly greater than 0.
	for counter > 0 {

		// Display current countdown second.
		fmt.Printf("T-minus %d second(s)...\n", counter)

		// Optional: Add a 1-second delay for realistic launch timing simulation.
		time.Sleep(1 * time.Second)

		// Decrement counter variable (equivalent to counter = counter - 1).
		counter--
	}

	// Output terminal launch message once the loop condition turns false.
	fmt.Println("🚀 Liftoff! We have ignition!")
}
