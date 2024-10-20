# Two Sum Problem

Given an array of integers nums and an integer target, return the indices of the two numbers such that they add up to the target.
You cannot use the same element twice, and you may assume that there is exactly one solution.

```
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```
| **Solution**          | **Time Complexity**       | **Space Complexity** | **Use Case**                                    |
|-----------------------|----------------------------|-----------------------|-------------------------------------------------|
| **Brute Force**       | \(O(n^2)\)                 | \(O(1)\)              | Works for small arrays; not efficient for large inputs |
| **Two-Pass Hash Map** | \(O(n)\)                   | \(O(n)\)              | Ideal for beginners or when clarity is needed        |
| **Single-Pass Hash Map** | \(O(n)\)                | \(O(n)\)              | More efficient; best for larger arrays               |
| **Two Pointers**      | \(O(n)\) / \(O(n \log n)\)*| \(O(1)\) / \(O(n)\)*  | Only for sorted arrays; may require sorting         |

