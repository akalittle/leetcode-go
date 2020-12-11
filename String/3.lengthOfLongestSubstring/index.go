package String

//Given a string s, find the length of the longest substring without repeating characters.

//Example 1:
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//Example 2:
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//Example 4:
//
//Input: s = ""
//Output: 0

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0

	window := make(map[byte]int)

	res := 0

	for right < len(s) {
		c := s[right]
		right++

		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		res = max(res, right-left)
	}

	return res
}
