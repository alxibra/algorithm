package main

import "fmt"

/*
Given a string, reverse it in-place using the two-pointer approach.
For example, if the input is "hello", the output should be "olleh".
*/

func main() {
	// bruteForce("olleh")
	twoPointer("olleh")
}

func twoPointer(word string) {
	left, right := 0, len(word)-1
	reserved := make([]byte, len(word))

	for left <= right {
		reserved[right], reserved[left] = word[left], word[right]
		left++
		right--
	}
	fmt.Println(string(reserved))
}

func bruteForce(word string) {
	reverse := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		j := len(word) - 1 - i
		reverse[i] = word[j]
	}
	fmt.Println(string(reverse))
}
