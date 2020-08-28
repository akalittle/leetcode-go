package Math

//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output: 321
//Example 2:
//
//Input: -123
//Output: -321
//Example 3:
//
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,   231 − 1].For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

func reverse(x int) int {
	res := 0

	for x != 0 {
		//  判断是否溢出 temp * 10 如果溢出的话 再 / 10 不会等于temp
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}

		res = res*10 + x%10
		x = x / 10
	}

	return res
}
