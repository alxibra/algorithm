package main

import (
	"fmt"
	"math"
)

func main() {
	// Example input
	nums := []int{1, 3, 5, 7, 10}
	target := 7
	// result := bruteForce(nums, target)
	result := binarySearch(nums, target)
	fmt.Println("Indices: ", result)
	fmt.Printf("elements: %d, %d\n", nums[result[0]], nums[result[1]])
}

func binarySearch(nums []int, target int) []int {
	n := len(nums)
	closestSum := math.MaxInt32
	closestIndices := []int{-1, -1}

	// n-1 not till last element
	for i := 0; i < n-1; i++ {
		complement := target - nums[i]
		j := binarySearchClosest(nums, i+1, n-1, complement)
		if j != -1 {
			currentSum := nums[i] + nums[j]
			if abs(target-currentSum) < abs(target-closestSum) {
				closestSum = currentSum
				closestIndices = []int{i, j}
			}
		}
	}

	return closestIndices
}

func binarySearchClosest(nums []int, left, right int, target int) int {
	closest := -1
	for left <= right {
		mid := left + (right-left)/2
		if closest == -1 || abs(nums[mid]) < abs(nums[closest]) {
			closest = mid
		}

		if nums[mid] < target {
			left = mid + 1
			continue
		}

		right = mid - 1
	}
	return closest
}

func bruteForce(nums []int, target int) []int {
	closestSum := math.MaxInt
	var closestIndeces []int
	fmt.Println(nums)

	// len(nums)-1 because we dont want to check until the end
	for i := 0; i < len(nums)-1; i++ {
		// len(nums) because we check until the end
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			fmt.Printf("sum: %d, i: %d, j: %d, nums[i]: %d, nums[j]: %d\n", sum, i, j, nums[i], nums[j])
			if abs(target-sum) < abs(target-closestSum) {
				closestSum = sum
				closestIndeces = []int{i, j}
			}
		}
		fmt.Println("...")
	}

	return closestIndeces

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}