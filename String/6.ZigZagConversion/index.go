package String

import "strings"

//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//And then read line by line: "PAHNAPLSIIGYIR"
//
//Write the code that will take a string and make this conversion given a number of rows:
//
//string convert(string s, int numRows)
//Example 1:
//
//Input: s = "PAYPALISHIRING", numRows = 3
//Output: "PAHNAPLSIIGYIR"
//Example 2:
//
//Input: s = "PAYPALISHIRING", numRows =  4
//Output: "PINALSIGYAHRPI"
//Explanation:
//
//P     I    N
//A   L S  I G
//Y A   H R
//P     I

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var b = []rune(s)
	var res = make([]string, numRows)
	var length = len(b)
	var period = numRows*2 - 2
	for i := 0; i < length; i++ {
		var mod = i % period
		if mod < numRows {
			res[mod] += string(b[i])
		} else {
			res[period-mod] += string(b[i])
		}
	}
	return strings.Join(res, "")

}
