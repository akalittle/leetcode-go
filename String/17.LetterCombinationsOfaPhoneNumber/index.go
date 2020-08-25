package String

//Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below.Note that 1 does not map to any letters.
//
//
//Example:
//
//Input: "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//Note:
//
//Although the above answer is in lexicographical order, your answer could be in any order you want.

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	result := []string{""}

	for i := 0; i < len(digits); i++ {
		t := m[digits[i]-'0']
		temp := []string{}
		for j := 0; j < len(t); j++ {
			for z := 0; z < len(result); z++ {
				temp = append(temp, result[z]+string(t[j]))
			}
		}
		result = temp
	}

	return result

}
