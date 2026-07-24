/*
Exercise 3: Letter Grade Converter

Write a program that converts a numerical grade (0-100) into a letter grade.

Requirements:
- Define an integer variable representing a student's score.
- Convert the score to a letter grade based on the following scale:
  * A: score >= 90
  * B: score >= 80
  * C: score >= 70
  * D: score >= 60
  * F: score < 60
- Print a formatted message detailing both the numerical score and the assigned letter grade.
*/

package main

import "fmt"

func main() {
	// Numerical score to evaluate.
	studentScore := 65

	// Evaluate the score from highest to lowest threshold.
	// Evaluating in descending order ensures higher thresholds take priority
	// without needing complex range conditions (e.g., studentScore >= 80 && studentScore < 90).
	if studentScore >= 90 {
		fmt.Printf("Score: %d | Final Grade: A\n", studentScore)
	} else if studentScore >= 80 {
		fmt.Printf("Score: %d | Final Grade: B\n", studentScore)
	} else if studentScore >= 70 {
		fmt.Printf("Score: %d | Final Grade: C\n", studentScore)
	} else if studentScore >= 60 {
		fmt.Printf("Score: %d | Final Grade: D\n", studentScore)
	} else {
		fmt.Printf("Score: %d | Final Grade: F\n", studentScore)
	}
}
