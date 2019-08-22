package solutions

/*

***
Runtime: 4 ms
Memory Usage: 3.8 MB
***

-- TwoSum(nums []int, target int) []int --
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

//TwoSum determines if two values equal the target and return their indices
// Successful whether or not the slice is sorted
func TwoSum(nums []int, target int) []int {
	var sln []int
	m := make(map[int]int)
	for i, n := range nums {
		if compl, ok := m[target-n]; ok {
			sln = append(sln, compl, i)
		} else {
			m[n] = i
		}
	}
	if len(sln) == 0 {
		sln = append(sln, 0, 0)
	}
	return sln
}
