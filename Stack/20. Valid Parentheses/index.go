package Stack

func isValid(s string) bool {
	brackets := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	var stack []rune

	for _, str := range s {
		if str == '{' || str == '[' || str == '(' {
			stack = append(stack, str)
		} else if len(stack) > 0 && brackets[str] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
