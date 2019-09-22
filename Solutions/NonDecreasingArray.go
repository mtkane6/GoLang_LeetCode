package solutions

/*

Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

Example 1:
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
Example 2:
Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.
Note: The n belongs to [1, 10,000].

----------

Runtime: 24 ms, faster than 86.96% of Go online submissions for Non-decreasing Array.
Memory Usage: 6.2 MB, less than 100.00% of Go online submissions for Non-decreasing Array.

*/

// CheckPossibility checks if array can be made nondecreasing by changing at most 1 value
func CheckPossibility(nums []int) bool {
	outOfOrderCount := 0
	highest := 0
	highestIndex := 0

	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			outOfOrderCount++
			if outOfOrderCount > 1 {
				return false
			}
			if i != 0 && i != len(nums)-2 && nums[i-1] > nums[i+1] && nums[i] > nums[i+2] {
				return false
			}
		}
		if nums[i] >= highest {
			highest = nums[i]
			highestIndex = i
		}
		if i == len(nums)-2 {
			highest = nums[i+1]
		}
	}

	if highestIndex == len(nums)-2 {
		return true
	}

	if highestIndex != 0 && (nums[highestIndex-1] > nums[highestIndex+1]) {
		if highestIndex < len(nums)-1 && nums[highestIndex] <= nums[highestIndex+2] {
			return true
		}
		return false
	}

	return true
}
