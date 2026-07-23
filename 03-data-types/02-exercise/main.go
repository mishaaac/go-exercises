/*
Exercise 2: Bank Account Struct & Zero-Value Initialization

Create a data structure (`struct`) representing a Bank Account storing:
- User ID (`int64` / long integer)
- Current balance (`float64` / decimal)
- Locked status (`bool` / boolean)

Initialize the account using Go's default zero values, inspect whether the account
defaults to locked or unlocked, and verify uninitialized user state checks.
*/

package main

import "fmt"

// BankAccount defines the schema for a bank account domain entity.
type BankAccount struct {
	ID       int64   // Defaults to 0
	Balance  float64 // Defaults to 0.0
	IsLocked bool    // Defaults to false
}

func main() {
	// Declare an uninitialized struct variable.
	// Go automatically assigns default zero values to all fields:
	// ID: 0, Balance: 0.0, IsLocked: false
	var account BankAccount

	fmt.Println("======= Bank Account Default State =======")

	// Check default lock status
	if account.IsLocked {
		fmt.Println("🔒 Account Status: Locked for transactions")
	} else {
		fmt.Println("🔓 Account Status: Unlocked (Ready for transactions)")
	}

	// Verify uninitialized registration state
	if account.ID == 0 {
		fmt.Println("⚠️  Alert: User registration incomplete (Unassigned ID).")
	}

	fmt.Println("\n======= Internal Struct Representation =======")
	// %#v prints full struct field names alongside their assigned values
	fmt.Printf("%#v\n", account)
}