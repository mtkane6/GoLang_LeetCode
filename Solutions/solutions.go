package solutions

// LongestCommonPrefix returns the longest common prefix string
func LongestCommonPrefix(strs []string) string {
	Prefix := ""
	if len(strs) > 0 {
		for i := range strs[0] {
			for _, curr := range strs {
				if len(curr) == i || curr[i] != strs[0][i] {
					return Prefix
				}
			}
			Prefix = Prefix + string(strs[0][i])
		}
	}
	return Prefix
}

//TwoSum determines if two values equal the target and return their indices
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

// IsValid Valid Parentheses ==== 100% faster than all other Go submissions, 50% smaller (2mb)
func IsValid(s string) bool {
	var stack = Stack{
		top: nil,
		len: 0,
	}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := range s {
		if stack.top != nil && stack.top.val == m[s[i]] && len(m) != 0 {
			if stack.len > 0 {
				stack.top = stack.top.prev
			}
			stack.len--
		} else if _, ok := m[s[i]]; ok {
			return false
		} else { // add to map and to stack
			stack.top = &node{
				val:  s[i],
				prev: stack.top,
			}
			stack.len++
		}
	}
	if stack.len == 0 {
		return true
	}
	return false
}

// Stack is the stack struct
type Stack struct {
	top *node
	len int
}

type node struct {
	val  byte
	prev *node
}
