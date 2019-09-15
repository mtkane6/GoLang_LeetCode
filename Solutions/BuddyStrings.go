package solutions

/*
Example 1:

Input: A = "ab", B = "ba"
Output: true
Example 2:

Input: A = "ab", B = "ab"
Output: false
Example 3:

Input: A = "aa", B = "aa"
Output: true
Example 4:

Input: A = "aaaaaaabc", B = "aaaaaaacb"
Output: true
Example 5:

Input: A = "", B = "aa"
Output: false

---------------------------------------------
"Runtime: 0 ms, faster than 100.00% of Go online submissions for Buddy Strings.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Buddy Strings."
*/

// BuddyStrings returns true if 2 letters in A can be swapped to make it equal to B.
func BuddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if len(A) < 2 {
		return false
	}

	if A == B {
		m := make(map[byte]int)
		for i := range A {
			temp := m[A[i]]
			m[A[i]] = temp + 1
			if m[A[i]] == 2 {
				return true
			}
		}
	}

	miss1, miss2, count := -1, -1, 0

	for i := range A {
		if A[i] != B[i] {
			count++
			if count > 2 {
				return false
			}

			miss1 = miss2
			miss2 = i
		}
	}
	if count != 2 || A[miss1] != B[miss2] || A[miss2] != B[miss1] {
		return false
	}

	return true
}
