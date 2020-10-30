package LCOF

//找出数组中重复的数字。
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
//示例 1：
//
//输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3

func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			swap(i, nums[i], nums)
		}
	}

	return -1
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
