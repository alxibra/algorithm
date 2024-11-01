package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println("count: ", bruteForce(nums, k))
}

func bruteForce(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for l := i; l < len(nums); l++ {
			sum += nums[l]
			if sum == k {
				count++
			}
		}
	}
	return count
}
