package main

import (
	"fmt"
	"math"
)

func main() {
	// Example input
	nums := []int{1, 3, 5, 7, 10}
	target := 7
	result := bruteForce(nums, target)
	fmt.Println("Indeces: ", result)
	fmt.Printf("elements: %d, %d\n", nums[result[0]], nums[result[1]])
}

func bruteForce(nums []int, target int) []int {
	closestSum := math.MaxInt
	var closestIndeces []int

	// len(nums)-1 because we dont want to check until the end
	for i := 0; i < len(nums)-1; i++ {
		// len(nums) because we check until the end
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if abs(target-sum) < abs(target-closestSum) {
				closestSum = sum
				closestIndeces = []int{i, j}
			}
		}
	}

	return closestIndeces

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
