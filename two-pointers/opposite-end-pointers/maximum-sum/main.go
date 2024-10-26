package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 6, 3, 5, 7}
	bruteForce(nums)
	singlePass(nums)
}

func singlePass(nums []int) {
	maxFirst, maxSecond := math.MinInt32, math.MinInt32
	for _, v := range nums {
		if v > maxFirst {
			maxSecond = maxFirst
			maxFirst = v
			continue
		}
		if v > maxSecond {
			maxSecond = v
			continue
		}
	}
	sum := maxFirst + maxSecond

	fmt.Println("max: ", sum)
}

func bruteForce(nums []int) {
	maxSum := math.MinInt32
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if maxSum < sum {
				maxSum = sum
			}
		}
	}
	fmt.Println("max sum: ", maxSum)
}
