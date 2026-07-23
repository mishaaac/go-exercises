/*
Exercise 2: Variable Shadowing and Scope Correction

Correct the provided code so that it compiles and updates variables as intended,
explaining why the original code failed or produced unexpected results.

Original Issue:
The inner `attempts := 2` statement inside the `if` block used short variable
declaration (`:=`), which created a NEW, separate local variable that SHADOWED
the outer `attempts` variable! When the `if` block ended, the shadowed variable
was destroyed, leaving the outer `attempts` variable unchanged at 3.
*/

package main

import "fmt"

var globalMessage = "System startup"

func main() {
	// Outer scope variable declaration
	attempts := 3

	if attempts > 0 {
		// FIX: Use plain assignment '=' instead of short declaration ':='.
		// This modifies the existing 'attempts' variable in the outer scope
		// rather than creating a shadowed variable restricted to this block.
		attempts = 2
	}

	fmt.Println(globalMessage)
	fmt.Println("Remaining attempts:", attempts) // Outputs: 2
}
