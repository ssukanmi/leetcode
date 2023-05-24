package solutions

import "strconv"

// Given an integer x, return true if x is a palindrome, and false otherwise.

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints:
// -231 <= x <= 231 - 1

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverseInt(x)
}

func reverseInt(i int) int {
	reversed := 0
	for i != 0 {
		reversed = reversed*10 + i%10
		i /= 10
	}
	return reversed
}

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return s == reverseString(s)
}

func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, val := range s {
		n--
		runes[n] = val
	}
	return string(runes)
}
