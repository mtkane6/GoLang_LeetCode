package main

import (
	// s "LEETCODE/solutions"
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

	// t := []int{2, 4, 5, 1, 2}
	// test := s.CheckPossibility(t)
	// fmt.Println(test)

	// t := []int{0, 4, 5, 1, 2}
	// water := s.WaterPlants(t, 5)
	// fmt.Println(water)

	// A := []int{1, 2, 2, 2}
	// B := []int{2, 2, 3, 4}
	// dominos := s.MinDominoRotations(A, B)
	// fmt.Println(dominos)

	// 3, 6, 17, 0, 5, 4, 7, 9, 12, 1

	// fmt.Println("Quicksort: ", s.Quicksort(a))
	// fmt.Println("Mergesort: ", s.MergeSort(a))

	// b := make([]int, 5)
	// fmt.Println("b: ", b)
	// fmt.Println("cap: ", cap(b))
	// b = append(b, 0)
	// fmt.Println("b: ", b)
	// fmt.Println("cap: ", cap(b))
	// b = append(b, 0, 0)
	// fmt.Println("b: ", b)
	// fmt.Println("cap: ", cap(b))

	// var intSlice = new([]int)

	// fmt.Println(reflect.ValueOf(intSlice).Kind())
	// fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice)
	// if intSlice == nil {
	// 	fmt.Println("intSlice is nil")
	// }

	// val := s.HasGroupsSizeX(a)
	// fmt.Println(val)

	// fmt.Println(s.Gcd(10,15))

	I := 121
	// fmt.Println(s.Reverse(I))
	fmt.Println(s.IsPalindrome(I))
}
