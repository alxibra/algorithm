package main

import (
	"fmt"
)

func main() {
	word := []string{"b", "a", "c", "a", "b"}
	// result := bruteForce(word)
	result := twoPointers(word)
	fmt.Println("palindrom: ", result)
}

func bruteForce(word []string) bool {
	for i := 0; i < len(word)/2; i++ {
		left, right := word[i], word[len(word)-1-i]
		if left != right {
			return false
		}
	}
	return true
}

func twoPointers(word []string) bool {
	left, right := 0, len(word)-1
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}
