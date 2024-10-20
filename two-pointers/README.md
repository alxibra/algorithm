# Two Pointers Technique
**The Two Pointers** technique is a powerful and versatile approach used to solve various algorithmic problems,
particularly in **arrays and strings**.
The idea behind two pointers is to use two indices (or pointers) to traverse through a data structure simultaneously,
often reducing the need for nested loops and optimizing the time complexity of the problem.

## Core Concept
In the Two Pointers technique,
two pointers are used to traverse the array or string in a coordinated manner.
Depending on the problem, the two pointers can either:
1. Start at opposite ends and move towards each other (common in sorted arrays).
2. Start at the same point and move at different speeds (useful in linked lists and cycle detection problems).
3. Move within a sliding window to track subarrays or substrings.

## Key Principles:
1. 	Pointer Initialization:
    Typically, one pointer starts at the beginning (left),
    and the other pointer starts at the end (right).
    In some cases, both pointers may start at the same position.
2.  Pointer Movement:
    The pointers are moved according to the problem’s requirements.
    For example, move the left pointer rightward to increase values or
    the right pointer leftward to decrease values.
3.  Condition Checking:
    At each step, check the condition (e.g., sum of two numbers, checking for a palindrome)
    and adjust the pointers based on whether the condition is met.

## When to Use the Two-Pointer Technique
1.  **Sorted Data**:
    If the problem involves a sorted array or string,
    the two-pointer technique is often the best approach,
    especially for finding pairs or subarrays that meet specific conditions.
2.  **Looking for Pairs or Triplets**:
    If the problem involves finding two or three elements in an array that satisfy
    a condition (like spesific sum), two pointers can reduce the need for nested loops
3.  *Link List Problems*:
    When working with linked lists, two pointers can be used to detect cycles or
    find the middle of the list efficiently.
