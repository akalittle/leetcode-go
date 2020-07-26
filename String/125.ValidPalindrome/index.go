package String

import "strings"

//
//Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
//
//Note: For the purpose of this problem, we define empty string as valid palindrome.
//
//Example 1:
//
//Input: "A man, a plan, a canal: Panama"
//Output: true
//Example 2:
//
//Input: "race a car"
//Output: false

func isValidByte(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func isPalindrome(s string) bool {
	var s1 string

	lens := len(s)

	for i := 0; i < lens; i++ {
		if isValidByte(s[i]) {
			s1 += string(s[i])
		}
	}

	s1 = strings.ToLower(s1)
	n := len(s1)

	for i := 0; i < n/2; i++ {
		if s1[i] != s1[n-i-1] {
			return false
		}
	}

	return true
}
