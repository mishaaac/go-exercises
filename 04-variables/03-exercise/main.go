/*
Exercise 3: Bank Vault Access Control & Formal Declarations

Requirements:
- Declare globally allowed attempts (3) using formal syntax (`var maxAttempts int = 3`).
- Inside validateAccess(), declare a local variable for current failed attempts.
- Mutate a local boolean variable (`systemLocked`) to true if failed attempts >= maxAttempts.
- Ensure all declared variables are read and printed to avoid unused variable compiler errors.
*/

package main

import "fmt"

// Package-level global declaration using formal 'var' syntax with explicit type.
var maxAttempts int = 3

func validateAccess() {
	// Local variable tracking current failed attempts.
	failedAttempts := 3

	// System security lockout status flag.
	systemLocked := false

	// Lock system if failed attempts reach or exceed max allowed threshold.
	if failedAttempts >= maxAttempts {
		systemLocked = true
	}

	// Compliance with restriction: All declared variables are printed to standard output.
	fmt.Println("=== 🏦 Bank Vault Security Status ===")
	fmt.Printf("Max attempts: %d\n", maxAttempts)
	fmt.Printf("Failed attempts: %d\n", failedAttempts)
	fmt.Printf("System locked: %t\n", systemLocked)

	if systemLocked {
		fmt.Println("🚨 SECURITY ALERT: Vault access is permanently locked!")
	} else {
		fmt.Println("🔓 STATUS NORMAL: Remaining attempts available.")
	}
}

func main() {
	validateAccess()
}
