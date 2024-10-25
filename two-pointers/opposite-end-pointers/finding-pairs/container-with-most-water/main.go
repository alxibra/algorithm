package main

import (
	"fmt"
	"math"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := bruteForce(height)

	fmt.Println("result: ", result)
}

func bruteForce(height []int) int {
	maxCapacity := math.MinInt32
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			currentCapacity := math.Min(float64(height[i]), float64(height[j])) * float64(j-i)

			if currentCapacity > float64(maxCapacity) {
				maxCapacity = int(currentCapacity)
			}
		}
		fmt.Println("max capacity: ", maxCapacity)
	}
	return maxCapacity
}

func twoPointers(height []int) int {
	left, right := 0, len(height)-1
	maxCapacity := math.MinInt32
	for left < right {
		currentCapacity := math.Min(float64(height[left]), float64(height[right])) * float64(right-left)
		if currentCapacity > float64(maxCapacity) {
			maxCapacity = int(currentCapacity)
		}

		if height[left] < height[right] {
			left++
			continue
		}
		right--
	}
	return maxCapacity
}
