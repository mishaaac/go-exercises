/*
Exercise 2: Interactive Number Accumulator

Write an interactive program that continuously prompts the user to enter integers.
The program keeps a running total sum of all non-negative numbers entered.
The loop must terminate immediately when the user enters any negative number,
displaying the final total sum accumulated up to that point.

Requirements:
- Use an infinite loop (`for {}`) to process inputs continuously.
- Read integer input from standard input using `fmt.Scanln`.
- Handle scan parsing errors gracefully without crashing or panicking (`panic`).
- Check if the input is negative (`num < 0`); if so, exit the loop using `break`.
- Accumulate valid positive numbers into a running total using `+=`.
- Print the final total sum upon loop exit.
*/

package main

import "fmt"

func main() {
	// Running total sum of entered numbers.
	totalSum := 0

	fmt.Println("===== Number Accumulator =====")
	fmt.Println("Enter positive integers to add to the sum.")
	fmt.Println("Enter a negative number (e.g., -1) to stop and print the total.")

	// Infinite loop keeps processing inputs until a negative number is entered.
	for {
		var inputNumber int

		fmt.Print("Enter a number: ")
		_, err := fmt.Scanln(&inputNumber)

		// Graceful error handling: recover from invalid (non-integer) input
		// instead of crashing the program with panic().
		if err != nil {
			fmt.Println("[Error] Invalid input. Please enter a valid integer.")

			// Clear standard input buffer to prevent an infinite error loop.
			var dump string
			fmt.Scanln(&dump)
			continue
		}

		// Termination condition: Stop immediately on negative numbers.
		if inputNumber < 0 {
			fmt.Println("\nNegative number detected. Stopping accumulator.")
			break
		}

		// Accumulate non-negative number into running total.
		totalSum += inputNumber
		fmt.Printf("-> Added %d. Current running total: %d\n\n", inputNumber, totalSum)
	}

	// Output final accumulated result.
	fmt.Println("==============================")
	fmt.Printf("Final Accumulated Total: %d\n", totalSum)
	fmt.Println("==============================")
}
