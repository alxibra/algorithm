package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 6, 3, 5, 7}
	bruteForce(nums)
	singlePass(nums)
	twoPass(nums)
	Heap(nums)
}

func twoPass(nums []int) {
	maxFirst := math.MinInt32

	for _, v := range nums {
		if v > maxFirst {
			maxFirst = v
		}
	}

	maxSecond := math.MinInt32
	for _, num := range nums {
		if num > maxSecond && num < maxFirst {
			maxSecond = num
		}
	}
	sum := maxFirst + maxSecond
	fmt.Println("max: ", sum)
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(num any) {
	*h = append(*h, num.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func Heap(nums []int) {
	h := &MinHeap{}
	heap.Init(h)
	for index := 0; index < 2; index++ {
		h.Push(nums[index])
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	fmt.Println("heap: ", h)
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
