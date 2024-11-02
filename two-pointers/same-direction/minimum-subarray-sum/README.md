# Minimum Size Subarray Sum

Given an array of positive integers nums and an integer target, find the minimal length of a contiguous subarray [nums[l], nums[l+1], ..., nums[r-1], nums[r]] of which the sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1
```
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
```
Example 2

```
Input: target = 4, nums = [1,4,4]
Output: 1
Explanation: The subarray [4] meets the requirement and has length 1.
```
Example 3
```
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
Explanation: No subarray meets the requirement.
```

### Constraint
	•	1 <= target <= 10^9
	•	1 <= nums.length <= 10^5
	•	1 <= nums[i] <= 10^4

### Notes
	•	The elements in nums are positive integers.
	•	The goal is to find the shortest subarray with a sum that is at least target.
	•	If no such subarray exists, return 0.

