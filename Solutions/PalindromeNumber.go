package solutions

/*
Runtime: 16 ms, faster than 65.03% of Go online submissions for Palindrome Number.
Memory Usage: 5.2 MB, less than 62.50% of Go online submissions for Palindrome Number.

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

import (
	"strconv"
)

// IsPalindrome tests if a nuber is read the same backwards as it is forward.
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x <= 9 { // single digit is a palindrome
		return true
	}
	var ints []byte
	xString := strconv.Itoa(x)
	ints = []byte(xString)
	for i := 0; i <= (len(ints)-1)/2; i++ {
		if ints[i] != ints[(len(ints)-1)-i] {
			return false
		}
	}
	return true
}
