package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLength := math.MaxInt32 // Use a large number to represent the minimum length initially
	currentSum := 0            // Sum of the current window
	left := 0                  // Left pointer

	for right := 0; right < n; right++ {
		currentSum += nums[right] // Expand the window by adding nums[right]

		// Shrink the window as small as possible while the sum is >= target
		for currentSum >= target {
			// Update the minimum length if the current window is smaller
			minLength = min(minLength, right-left+1)
			currentSum -= nums[left] // Shrink the window by removing nums[left]
			left++                   // Move the left pointer to the right
		}
	}

	// If minLength was never updated, return 0 (no valid subarray found)
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test case
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println("Minimum subarray length:", minSubArrayLen(target, nums)) // Output: 2
}
