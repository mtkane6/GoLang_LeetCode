package solutions

import (
	"math"
)

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
Memory Usage: 2.2 MB, less than 80.00% of Go online submissions for Reverse Integer.

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit
signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your
function returns 0 when the reversed integer overflows.

*/

// Reverse reverses an input integer
func Reverse(x int) int {
	if x == 0 || x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	neg := false
	if x < 0 {
		x *= -1
		neg = true
	}
	y := 0
	for ; x > 0; x = x / 10 {
		y = (y*10 + x%10)
		if y > math.MaxInt32 {
			return 0
		}
	}
	if neg {
		y *= -1
	}
	return y
}
