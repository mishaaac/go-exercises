/*
Exercise 2: Age Stage Classifier

Write a program that classifies a person's age into different life stages.

Requirements:
- Define an age variable representing a person's age in years.
- Evaluate the age using a logical branching structure.
- Categorize the age into one of four stages:
  * "Child" for ages strictly under 13.
  * "Teenager" for ages between 13 and 19 (inclusive).
  * "Adult" for ages between 20 and 64 (inclusive).
  * "Senior" for ages 65 and older.
- Print a formatted, contextual message outputting the age and the corresponding stage.
*/

package main

import "fmt"

func main() {
	// Target age to classify in years.
	personAge := 23

	// Classify the stage of life.
	// Since each condition is evaluated sequentially, earlier conditions naturally
	// filter out lower bounds, allowing concise upper-bound checks.
	if personAge < 13 {
		fmt.Printf("Age %d is classified as: Child\n", personAge)
	} else if personAge <= 19 {
		// Evaluated only if personAge >= 13.
		fmt.Printf("Age %d is classified as: Teenager\n", personAge)
	} else if personAge <= 64 {
		// Evaluated only if personAge >= 20.
		fmt.Printf("Age %d is classified as: Adult\n", personAge)
	} else {
		// Evaluated only if personAge >= 65.
		fmt.Printf("Age %d is classified as: Senior\n", personAge)
	}
}
