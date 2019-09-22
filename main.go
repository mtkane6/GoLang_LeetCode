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
	// var a = []string{
	// 	"aa",
	// 	"a",
	// }
	// t := s.LongestCommonPrefix(a)
	// fmt.Println(t)

	// MergeTwoLists
	// var l1 = s.ListNode
	// {
	// 	Val: 1,
	// 	Next: &s.ListNode{
	// 		Val: 2,
	// 		Next: &s.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }

	// var l2 = s.ListNode{
	// 	Val: 1,
	// 	Next: &s.ListNode{
	// 		Val: 3,
	// 		Next: &s.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }
	// result := s.MergeTwoLists(&l1, &l2)
	// fmt.Printf("%+v", result)

	// BuddyStrings
	// A := "abab"
	// B := "abab"
	// result := s.BuddyStrings(A, B)
	// fmt.Println(result)

	// m := make(map[string]string)
	// m["first"] = "mitch"
	// m["second"] = "kane"

	// for k := range m {
	// 	fmt.Println(k)
	// }
	// for l := range m {
	// 	fmt.Println(m[l])
	// }

	//421
	//423
	//3412

	t := []int{0, 1, 2, 3, 4, 5, 7, 9, 6, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	test := s.CheckPossibility(t)
	fmt.Println(test)

}
