package String

import (
	"strconv"
)

//Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
//
//Note:
//
//The length of both num1 and num2 is < 5100.
//Both num1 and num2 contains only digits 0-9.
//Both num1 and num2 does not contain any leading zero.
//You must not use any built-in BigInteger library or convert the inputs to integer directly.
func addStrings(num1 string, num2 string) string {
	nums2Rune1 := []rune(num1)
	nums2Rune2 := []rune(num2)

	l1 := len(nums2Rune1) - 1
	l2 := len(nums2Rune2) - 1

	result := make([]rune, 0)

	carry := 0

	for l1 >= 0 || l2 >= 0 {
		v1 := 0
		v2 := 0

		if l1 >= 0 {
			v1 = int(nums2Rune1[l1] - '0')
		}

		if l2 >= 0 {
			v2 = int(nums2Rune2[l2] - '0')
		}

		value := v1 + v2 + carry

		carry = value / 10

		real := value % 10

		result = append([]rune(strconv.Itoa(real)), result...)

		l1--
		l2--

	}

	if carry > 0 {
		result = append([]rune("1"), result...)
	}

	return string(result)
}
