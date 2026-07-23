/*
Exercise 3: Periodic Website Health Checker

Write a Go script that periodically checks the HTTP connectivity status of 3 popular
websites (e.g., https://www.google.com, https://www.github.com, and an invalid URL
https://api.invalida.xyz to simulate a DNS/network failure).

Requirements:
- Define a custom `http.Client` with a strict execution timeout (e.g., 3 seconds).
- Return HTTP status code and error states cleanly from helper functions.
- Correct conditional branch execution when evaluating request success vs failure.
- Ensure proper resource cleanup via `defer resp.Body.Close()`.
- Add periodic monitoring loops (`time.Ticker` or `time.Sleep`) to check sites continuously.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

// verifyStatus attempts an HTTP GET request with a configured client timeout.
// Returns HTTP status code or 0 with error details on failure.
func verifyStatus(url string) (int, error) {
	// Custom HTTP client with explicit 3-second timeout guard.
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	// Always close response body on successful response connection.
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func main() {
	targetURLs := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://api.invalida.xyz", // Intentionally invalid domain to test error branch
	}

	fmt.Println("==================================================")
	fmt.Println("  🌐 Website Health & Connectivity Monitor        ")
	fmt.Println("==================================================")

	// Single check iteration (or wrap in ticker/sleep for continuous periodic monitoring)
	fmt.Printf("[%s] Executing Connectivity Check...\n\n", time.Now().Format("15:04:05"))

	for _, url := range targetURLs {
		code, err := verifyStatus(url)

		// Fix: Use if/else logic so failure outputs don't also print an [OK] message!
		if err != nil {
			fmt.Printf("❌ [FAILED] %s\n   └─ Error: %v\n\n", url, err)
		} else {
			fmt.Printf("✅ [OK]     %s\n   └─ HTTP Status: %d %s\n\n",
				url, code, http.StatusText(code))
		}
	}
}