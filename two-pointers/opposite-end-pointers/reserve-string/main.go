package main

import "fmt"

/*
Given a string, reverse it in-place using the two-pointer approach.
For example, if the input is "hello", the output should be "olleh".
*/

func main() {
	bruteForce("hello")
}

func bruteForce(word string) {
	reverse := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		j := len(word) - 1 - i
		reverse[i] = word[j]
	}
	fmt.Println(string(reverse))
}
