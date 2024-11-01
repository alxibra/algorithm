package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2}
	k := 2
	// fmt.Println("count: ", bruteForce(nums, k))
	fmt.Println("count: ", bruteForcePrefixSum(nums, k))
}

func bruteForce(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func bruteForcePrefixSum(nums []int, k int) int {
	n := len(nums)
	count := 0
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			sum := prefix[end]
			if start > 0 {
				sum -= prefix[start-1]
			}

			if sum == k {
				count++
			}
		}
	}
	return count
}
