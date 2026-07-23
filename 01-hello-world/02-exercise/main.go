/*
Exercise 2: Greeting Multiple Users with Constants

Write a program that greets three different users.
Define a constant at the beginning of the program with a basic welcome message
and combine it sequentially with each user's name.

Requirements:
- Define a constant string for the welcome message (`const welcomeMessage = "Welcome"`).
- Print customized welcome greetings for three different users.
- Use string formatting (`fmt.Printf`) or string concatenation to combine the constant and name.
*/

package main

import "fmt"

func main() {
	// Constant welcome prefix (untyped constant string).
	const welcomeMessage = "Welcome"

	// Output individual user greetings using string formatting (%s).
	fmt.Printf("%s, Cha!\n", welcomeMessage)
	fmt.Printf("%s, Marin!\n", welcomeMessage)
	fmt.Printf("%s, Reze!\n", welcomeMessage)
}
