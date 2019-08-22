package solutions

/*

***
Runtime: 0 ms
Memory Usage: 2.4 MB
***

-- LongestCommonPrefix(strs []string) string --
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

*/

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
