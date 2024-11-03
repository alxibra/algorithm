package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	minL := math.MaxInt32
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			minL = min(sum, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if minL == math.MaxInt32 {
		minL = 0
	}
	return minL
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bruteForce(k int, nums []int) int {
	minL := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum := 0
		for left := right; left < len(nums); left++ {
			sum += nums[left]
			if sum >= k {
				minL = min(minL, left-right+1)
				break
			}
		}
	}
	if minL == math.MaxInt32 {
		minL = 0
	}
	return minL
}

func main() {
	// Test case
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	// fmt.Println("Minimum subarray length:", minSubArrayLen(target, nums)) // Output: 2
	fmt.Println("Minimum subarray length:", bruteForce(target, nums)) // Output: 2
}
