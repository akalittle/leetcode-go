package String

import "math"

//Given two strings s and t, return the minimum window in s which will contain all the characters in t.If there is no such window in s that covers all characters in t, return the empty string "".
//
//Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.
//
//Example 1:
//
//Input: s = "ADOBECODEBANC", t = "ABC"
//Output: "BANC"
//Example 2:
//
//Input: s = "a", t = "a"
//Output: "a"
//
//
//Constraints:
//
//1 <= s.length, t.length <= 105
//s and t consist of English letters.

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	right := 0

	match := 0

	start, end := 0, 0
	min := math.MaxInt64

	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			window[c] ++
			if window[c] == need[c] {
				match++
			}
		}
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}

			c = s[left]
			left++
			if need[c] != 0 {
				if window[c] == need[c] {
					match--
				}
				window[c] --
			}
		}
	}

	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
