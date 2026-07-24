/*
Exercise 1: Number Sign Checker

Write a program that determines whether an integer is positive, negative, or zero.

Requirements:
- Define an integer variable.
- Evaluate whether the number is greater than zero, less than zero, or equal to zero.
- Print a clear, formatted message to the console displaying the evaluated number and its classification.
*/

package main

import "fmt"

func main() {
	// The target value to evaluate.
	targetNumber := -5

	// Evaluate the numeric sign using standard control flow constructs.
	// Since the conditions are mutually exclusive, chaining 'else if' ensures
	// only one execution path is taken.
	if targetNumber > 0 {
		fmt.Printf("The number %d is positive.\n", targetNumber)
	} else if targetNumber < 0 {
		fmt.Printf("The number %d is negative.\n", targetNumber)
	} else {
		fmt.Printf("The number %d is zero.\n", targetNumber)
	}
}
