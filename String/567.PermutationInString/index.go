package String

//Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
//
//Example 1:
//
//Input: s1 = "ab" s2 = "eidbaooo"
//Output: True
//Explanation: s2 contains one permutation of s1 ("ba").
//Example 2:
//
//Input:s1 = "ab" s2 = "eidboaoo"
//Output: False
// 
//
//Constraints:
//
//The input strings only contain lower case letters.
//The length of both given strings is in range [1, 10, 000].

func checkInclusion(s1 string, s2 string) bool {
	window := make(map[byte]int)
	need := make(map[byte]int)

	for i := range s1 {
		need[s1[i]]++
	}

	left := 0
	right := 0
	match := 0

	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] != 0 {
			window[c]++
			if window[c] == need[c] {
				match++
			}
		}

		for right-left >= len(s1) {
			if match == len(need) {
				return true
			}

			d := s2[left]
			left++

			if need[d] != 0 {
				if window[d] == need[d] {
					match--
				}
				window[d]--
			}
		}
	}

	return false
}
