package solutions

/*

***
Runtime: 28 ms
Memory Usage: 7.3 MB
***

-- FindTarget(root *TreeNode, k int) bool --
Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True


Example 2:

Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False

*/

// TreeNode is node for findTarget()
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findTarget(BST)
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int)
	buildMap(m, root)
	for key := range m {
		if _, ok := m[k-key]; ok {
			return true
		}
	}
	return false
}

func buildMap(m map[int]int, root *TreeNode) {
	if root != nil {
		m[root.Val] = 0
	}
	buildMap(m, root.Left)
	buildMap(m, root.Right)
}
