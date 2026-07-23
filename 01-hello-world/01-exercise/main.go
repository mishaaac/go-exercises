/*
Exercise 1: Basic Console Output

Write a Go program that displays three consecutive lines on the screen:
- Your name
- Your country of origin
- Your current favorite programming language

Requirements:
- Use the standard `fmt` package for terminal output.
- Output each piece of information on its own distinct line using `fmt.Println`.
*/

package main

import "fmt"

func main() {
	// Display profile information sequentially on separate lines.
	fmt.Println("Name: Cristhian Albites")
	fmt.Println("Country: Peru")
	fmt.Println("Favorite language: Go")
}
