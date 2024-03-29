package solutions

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.5 MB, less than 64.29% of Go online submissions for Merge Two Sorted Lists.
*/

// ListNode is the Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists merges two lists
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = MergeTwoLists(l1, l2.Next)
	return l2
}
