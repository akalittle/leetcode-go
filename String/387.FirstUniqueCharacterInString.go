package String


//Given a string, find the first non-repeating character in it and return its index.If it doesn't exist, return -1.
//
//Examples:
//
//s = "leetcode"
//return 0.
//
//s = "loveleetcode"
//return 2.
// 
//
//Note: You may assume the string contains only lowercase English letters.


func firstUniqChar(s string) int {
	var freq [26]int

	for _, v := range s {
		freq[v-'a']++
	}

	for k, v := range s {
		if freq[v-'a'] == 1 {
			return k
		}
	}

	return -1
}
