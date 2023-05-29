package solutions

import (
	"sort"
	"strings"
)

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.

func LongestCommonPrefix(strs []string) string {
	// initially set prefix to the first string
	pref := strs[0]
	// loop through all the strings in the array
	for _, str := range strs {
		// continuesly check if each string has prefix
		// if it doesn't reduce the prefix by the last letter
		for !strings.HasPrefix(str, pref) {
			pref = pref[:len(pref)-1]
		}
	}
	return pref
}

func LongestCommonPrefix2(strs []string) string {
	// if only one element return it
	if len(strs) == 1 {
		return strs[0]
	}

	// sort the strings array and compare the first with the last
	// they will be the most different
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]
	for idx := range first {
		if first[idx] != last[idx] {
			return first[:idx]
		}
	}
	return first
}
