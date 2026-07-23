/*
Exercise 2: Querying Public API Headers with net/http

Write a Go program that issues an HTTP GET request to the public API
"https://api.github.com" using Go's standard `net/http` package.
Extract and print only the "Content-Type" response header returned by the server.

Requirements:
- Import standard library packages: `fmt`, `log`, and `net/http`.
- Make an HTTP GET request using `http.Get()`.
- Handle network/HTTP errors gracefully without using `panic()`.
- Ensure resource cleanup by closing `resp.Body` using `defer resp.Body.Close()`.
- Access response headers using `resp.Header.Get("Content-Type")`.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	endpoint := "https://api.github.com"

	fmt.Printf("Fetching response headers from %s...\n\n", endpoint)

	// Issue HTTP GET request to GitHub public API.
	resp, err := http.Get(endpoint)
	if err != nil {
		// Log error and exit gracefully instead of panicking.
		log.Fatalf("[Error] Failed to reach endpoint: %v", err)
	}

	// Always close response body when done to release network sockets.
	defer resp.Body.Close()

	// Extract specific HTTP header using response header map getter.
	contentType := resp.Header.Get("Content-Type")

	// Print retrieved header result.
	fmt.Printf("HTTP Status Code: %d %s\n", resp.StatusCode, resp.Status)
	fmt.Printf("Content-Type Header: %s\n", contentType)
}
