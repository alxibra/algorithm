package main

import (
	"fmt"
	"math"
)

func main() {
	// s := "abcabcbb"
	// fmt.Println("longest: ", bruteForce(s))
	// s2 := "bbbbb"
	// fmt.Println("longest: ", bruteForce(s2))
	// s3 := "pwwkew"
	// fmt.Println("longest: ", bruteForce(s3))
	s := "aabbca"
	fmt.Println("longest: ", twoPointer(s))
	// s2 := "bbbbb"
	// fmt.Println("longest: ", twoPointer(s2))
	// s3 := "pwwkew"
	// fmt.Println("longest: ", twoPointer(s3))
}

func bruteForce(s string) int {
	longest := math.MinInt32
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isUniq(s, i, j) {
				if longest < (j - i + 1) {
					longest = j - i + 1
				}
			}
		}
	}
	return longest
}

// Two Pointers with Frequency Array solution
func twoPointer(s string) int {
	n := len(s)
	if n == 0 {
		return 0 // Edge case: Empty string
	}

	// Create a frequency array to track occurrences of characters
	freq := [128]int{} // Assuming standard ASCII character set

	// Initialize two pointers and max length
	left, maxLength := 0, 0

	// Iterate with the right pointer
	for right := 0; right < n; right++ {
		char := s[right]

		// Increment the frequency of the current character
		freq[char]++

		// If the current character is repeated, shrink the window from the left
		for freq[char] > 1 {
			freq[s[left]]-- // Decrement the frequency of the character at the left pointer
			left++          // Move the left pointer to the right
		}

		// Update max length if the current window is longer
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func isUniq(s string, i, k int) bool {
	h := make(map[byte]bool)
	for l := i; l <= k; l++ {
		if h[s[l]] {
			return false
		}
		h[s[l]] = true
	}
	return true
}
