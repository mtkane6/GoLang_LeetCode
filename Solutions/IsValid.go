package solutions

/*

-- IsValid --
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

*/

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
