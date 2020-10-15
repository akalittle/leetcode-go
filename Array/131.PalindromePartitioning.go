package Array


//Given a string s, partition s such that every substring of the partition is a palindrome.
//
//Return all possible palindrome partitioning of s.
//
//Example:
//
//Input: "aab"
//Output:
//[
//["aa", "b"],
//["a", "a", "b"]
//]



func isValid(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	res := [][]string{}
	var backtrace func(s string, path []string, pos int)
	backtrace = func(s string, path []string, pos int) {
		if pos == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := pos; i < len(s); i++ {
			if !isValid(s, pos, i) {
				continue
			}

			path = append(path, s[pos:i+1])
			backtrace(s, path, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrace(s, []string{}, 0)
	return res
}
