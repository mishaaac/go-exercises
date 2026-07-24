/*
Exercise 1: Time Duration Unit Converter

Write a program that converts a given duration in total minutes into hours
and remaining minutes using fundamental arithmetic operators.

Requirements:
- Define an integer variable 'totalMinutes' representing total elapsed time (e.g., 135 minutes).
- Calculate full hours using integer division ('/').
- Calculate remaining minutes using the modulus operator ('%').
- Output a clear, formatted summary of the converted time duration.
*/

package main

import "fmt"

func main() {
	// Total elapsed time in minutes.
	totalMinutes := 135

	// Number of minutes in one standard hour.
	const minutesPerHour = 60

	// Arithmetic Calculations:
	// 1. Integer division truncated toward zero yields full hours (135 / 60 = 2).
	hours := totalMinutes / minutesPerHour

	// 2. Modulus operator calculates the remainder after division (135 % 60 = 15).
	remainingMinutes := totalMinutes % minutesPerHour

	// Output formatted conversion report.
	fmt.Println("=== Time Duration Converter ===")
	fmt.Printf("Input Duration:     %d minutes\n", totalMinutes)
	fmt.Printf("Converted Duration: %d hours and %d minutes\n", hours, remainingMinutes)
}
