package main

import (
	"fmt"
	"math"
)

func main() {
	s := "abcabcbb"
	fmt.Println("longest: ", bruteForce(s))
	s2 := "bbbbb"
	fmt.Println("longest: ", bruteForce(s2))
	s3 := "pwwkew"
	fmt.Println("longest: ", bruteForce(s3))
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
