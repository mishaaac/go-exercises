/*
Exercise 1: Global vs Local Variables & Server Memory Metrics

Write a Go program that declares:
- A global (package-level) variable representing the server name ("Production-1").
- Two local variables within main() representing available RAM (16 GB) and used RAM (12 GB).
- Calculate and display the remaining free RAM memory.

Requirements:
- Declare package-level variable `ServerName`.
- Declare function-scoped local variables for total and used memory metrics.
- Perform subtraction arithmetic (`total - used`).
- Display formatted output with units (GB).
*/

package main

import "fmt"

// Package-level (global) variable accessible across the main package.
var ServerName string = "Production-1"

func main() {
	// Function-scoped local variables initialized with short variable syntax (:=).
	totalRAM := 16 // Total memory in GB
	usedRAM := 12  // Used memory in GB

	// Calculate remaining unallocated memory.
	freeRAM := totalRAM - usedRAM

	fmt.Println("=== Server Memory Diagnostics ===")
	fmt.Printf("Server Name : %s\n", ServerName)
	fmt.Printf("Total Memory: %d GB\n", totalRAM)
	fmt.Printf("Used Memory : %d GB\n", usedRAM)
	fmt.Printf("Free Memory : %d GB\n\n", freeRAM)

	fmt.Printf("Status: Server %s has %d GB free out of %d GB total.\n",
		ServerName, freeRAM, totalRAM)
}
