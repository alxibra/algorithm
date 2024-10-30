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
	fmt.Println("longest: ", asciiArray(s))
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
	max, p2 := 0, 0
	frq := [128]int{}

	for p1 := 0; p1 < len(s); p1++ {
		chr := s[p1]
		frq[chr]++

		for frq[chr] > 1 {
			frq[s[p2]]--
			p2++
		}

		lng := p1 - p2 + 1
		if lng > max {
			max = lng
		}
	}
	return max
}

func asciiArray(s string) int {
	max, left := 0, 0

	lastSeen := make([]int, 128)
	for i := 0; i < 128; i++ {
		lastSeen[i] = -1
	}

	for right := 0; right < len(s); right++ {
		char := s[right]

		if lastSeen[char] >= left {
			left = lastSeen[char] + 1
		}
		lastSeen[char] = right

		length := right - left + 1
		if length > max {
			max = length
		}
	}
	return max
}

func twoPointerWithHashmap(s string) int {
	max, left := 0, 0
	set := make(map[byte]bool)

	for right := 0; right < len(s); right++ {
		char := s[right]
		for set[char] {
			delete(set, s[left])
			left++
		}
		set[char] = true
		length := right - left + 1
		if length > max {
			max = length
		}
	}
	return max
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
