package String

//Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
//
//Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20, 100.
//
//The order of output does not matter.
//
//Example 1:
//
//Input:
//s: "cbaebabacd" p: "abc"
//
//Output:
//[0, 6]
//
//Explanation:
//The substring with start index = 0 is "cba", which is an anagram of "abc".
//The substring with start index = 6 is "bac", which is an anagram of "abc".
//Example 2:
//
//Input:
//s: "abab" p: "ab"
//
//Output:
//[0, 1, 2]
//
//Explanation:
//The substring with start index = 0 is "ab", which is an anagram of "ab".
//The substring with start index = 1 is "ba", which is an anagram of "ab".
//The substring with start index = 2 is "ab", which is an anagram of "ab".

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)

	for i := range p {
		need[p[i]] ++
	}

	window := make(map[byte]int)

	left := 0
	right := 0

	match := 0

	res := []int{}

	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			window[c] ++
			if window[c] == need[c] {
				match++
			}
		}

		if right-left >= len(p) {
			if match == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if need[d] != 0 {
				if window[d] == need[d] {
					match--
				}
				window[d]--
			}
		}

	}

	return res

}
