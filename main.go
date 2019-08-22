package main

import (
	s "LEETCODE/solutions"
	"fmt"
)

func main() {
	// two sums
	// ------------
	// a := []int{-3, 4, 3, 90}
	// f := s.TwoSum(a, 0)
	// fmt.Println(f)

	// findTarget(BST)
	// ------------

	// Valid Parentheses
	// ------------
	// t := s.IsValid("([{}]{})")
	// fmt.Println(t)

	// Longest Common Prefix
	var a = []string{
		"aa",
		"a",
	}
	t := s.LongestCommonPrefix(a)
	fmt.Println(t)
}
