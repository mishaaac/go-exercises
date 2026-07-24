/*
Exercise 2: Bitmask Permission Flags using Go 'iota'

Represent Unix-style file permissions (Read, Write, Execute) using bitmasks
generated with bitwise left-shift operations and 'iota'.

Requirements:
- Define a custom type 'Permission' based on 'int'.
- Use 'iota' inside a 'const' block to auto-generate power-of-two bit flags:
  * Read:    1 << 0 (binary 001 = 1)
  * Write:   1 << 1 (binary 010 = 2)
  * Execute: 1 << 2 (binary 100 = 4)
- Demonstrate bitmask evaluation and formatting using binary formatting verbs (%b).
*/

package main

import "fmt"

// Permission defines a custom type for bitwise flag masks.
type Permission int

// Bitmask constants representing POSIX-style file permissions.
// 'iota' auto-increments per line (0, 1, 2...), shifting bit '1' to create unique bit positions.
const (
	Read    Permission = 1 << iota // 1 << 0 = 00000001 (1)
	Write                          // 1 << 1 = 00000010 (2) — expression '1 << iota' auto-repeats
	Execute                        // 1 << 2 = 00000100 (4)
)

func main() {
	fmt.Println("=== File System Bitmask Permissions ===")

	// %03b pads output with leading zeroes across 3 binary positions.
	fmt.Printf("Read Permission:    %03b (decimal %d)\n", Read, Read)
	fmt.Printf("Write Permission:   %03b (decimal %d)\n", Write, Write)
	fmt.Printf("Execute Permission: %03b (decimal %d)\n", Execute, Execute)

	// Combine permissions using bitwise OR (|)
	userPerms := Read | Execute
	fmt.Println("---------------------------------------")
	fmt.Printf("Combined (Read+Exec): %03b (decimal %d)\n", userPerms, userPerms)

	// Test if Write permission is granted using bitwise AND (&)
	hasWrite := (userPerms & Write) != 0
	fmt.Printf("Has Write Access?     %t\n", hasWrite)
}
