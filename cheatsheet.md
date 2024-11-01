```go
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
```

** prefix Sum **

```go
func prefixSum(nums []int) []int {
    n := len(nums)
    prefix := make([]int, n)
    prefix[0] = nums[0]
    for i := 1; i < n; i++ {
        prefix[i] = prefix[i-1] + nums[i]
    }
    return prefix
}

func subarraySum(prefix []int, i, j int) int {
    if i == 0 {
        return prefix[j]
    }
    return prefix[j] - prefix[i-1]
}

func main() {
    nums := []int{1, 2, 3, 4}
    prefix := prefixSum(nums)
    fmt.Println("Sum of subarray nums[1:3]:", subarraySum(prefix, 1, 3)) // Output: 9 (2 + 3 + 4)
}
```
Output: For nums = [1, 2, 3, 4], the prefix sum array would be [1, 3, 6, 10].
