package _27_Remove_Element

func removeElement(nums []int, val int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	i, j := 0, 0

	for j < length {
		if nums[j] == val {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return length - (j - i)
}
