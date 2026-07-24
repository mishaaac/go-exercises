/*
Exercise 3: Safe Integer Addition & Overflow Protection

Design a safe arithmetic addition utility for high-reliability systems (e.g., banking/finance)
that detects and prevents 32-bit signed integer (int32) overflow and underflow errors.

Requirements:
- Define a function 'safeAddInt32(a, b int32) (int32, error)'.
- Protect arithmetic operations against both upper bound overflow (> math.MaxInt32)
  and lower bound underflow (< math.MinInt32).
- Return a descriptive error message when bounds are exceeded.
- Implement comprehensive table-driven tests verifying safe bounds and boundary edge-cases.
*/

package main

import (
	"errors"
	"fmt"
	"math"
)

// Sentinel error values for arithmetic boundary violations.
var (
	ErrInt32Overflow  = errors.New("arithmetic error: operation resulted in int32 overflow")
	ErrInt32Underflow = errors.New("arithmetic error: operation resulted in int32 underflow")
)

// safeAddInt32 performs addition of two int32 operands safely.
// Casts operands to 64-bit integers (int64) during calculation to detect
// 32-bit boundary overflow or underflow prior to downcasting.
func safeAddInt32(a, b int32) (int32, error) {
	// Widen operands to 64-bit signed integers to compute sum without risk of wrapping.
	sum64 := int64(a) + int64(b)

	// Validate upper boundary threshold (MaxInt32 = 2,147,483,647).
	if sum64 > math.MaxInt32 {
		return 0, fmt.Errorf("%w (value %d > %d)", ErrInt32Overflow, sum64, math.MaxInt32)
	}

	// Validate lower boundary threshold (MinInt32 = -2,147,483,648).
	if sum64 < math.MinInt32 {
		return 0, fmt.Errorf("%w (value %d < %d)", ErrInt32Underflow, sum64, math.MinInt32)
	}

	// Downcast safely back to int32 now that boundary compliance is guaranteed.
	return int32(sum64), nil
}

func main() {
	// Table-driven test suite testing standard calculations, limits, and boundary overflows.
	testCases := []struct {
		a, b int32
	}{
		{10, 20},
		{-10, -20},
		{math.MaxInt32, 0},
		{math.MinInt32, 0},
		{math.MaxInt32, 1},  // Triggers positive overflow
		{math.MinInt32, -1}, // Triggers negative underflow
		{1_000_000_000, 1_000_000_000},
		{math.MaxInt32, math.MaxInt32}, // Extreme upper bound overflow
	}

	fmt.Println("=== Safe Int32 Banking Addition Engine ===")

	for idx, tc := range testCases {
		result, err := safeAddInt32(tc.a, tc.b)
		if err != nil {
			fmt.Printf("Test #%d: %d + %d -> %v\n", idx+1, tc.a, tc.b, err)
		} else {
			fmt.Printf("Test #%d: %d + %d = %d\n", idx+1, tc.a, tc.b, result)
		}
	}
}
