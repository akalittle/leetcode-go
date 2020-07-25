package Array

func removeDuplicates(nums []int) int {
	length := len(nums)

	if length < 2 {
		return length
	}
	i, j := 1, 1

	val := nums[0]

	for j < length {
		if nums[j] == val {
			j++
		} else {
			val = nums[j]
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return length - (j - i)
}
