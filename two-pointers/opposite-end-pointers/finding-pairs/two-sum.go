package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	// result := twoSumTwoPointers(nums, target)
	// result := twoSumBruteForce(nums, target)
	// result := twoSumHashMap(nums, target)
	result := twoSumTwoPassHashMap(nums, target)
	if result == nil {
		fmt.Println("No pair is found")
		return
	}
	fmt.Println("Pair is found: ")
	fmt.Println("Index: ", result)
	fmt.Printf("elements: %d, %d\n", nums[result[0]], nums[result[1]])
}

func twoSumTwoPassHashMap(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}
	for i, v := range nums {
		complement := target - v
		index, found := mp[complement]
		if found {
			return []int{i, index}
		}
	}
	return nil
}

func twoSumHashMap(nums []int, target int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		index, exist := mp[complement]
		if exist {
			return []int{index, i}
		}
		mp[nums[i]] = i
	}
	return nil
}

func twoSumTwoPointers(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		}
		if sum < target {
			left++
			continue
		}
		right--
	}
	return nil
}

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
