package Array

import (
	"sort"
)

//Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
//Note:
//
//The solution set must not contain duplicate quadruplets.
//
//Example:
//
//Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
//A solution set is:
//[
//[-1, 0, 0, 1],
//[-2, -1, 1, 2],
//[-2, 0, 0, 2]
//]

func fourSum(nums []int, target int) [][]int {

	arrLen := len(nums)

	if arrLen < 4 {
		return nil
	}

	// 生序排序
	sort.Ints(nums)
	var equals [][]int
	for i := 0; i < arrLen-3; i++ {
		// 第一个元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < arrLen-2; j++ {
			// 第二个元素去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, l := j+1, arrLen-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				diff := target - sum
				if diff == 0 {
					equals = append(equals, []int{nums[i], nums[j], nums[k], nums[l]})
				}
				if diff > 0 {
					// 还差 第三个指针往右
					for cur := k; k < l && nums[k] == nums[cur]; k++ {
					}
				} else {
					// 大了  第四个指针 往左
					for cur := l; l > k && nums[l] == nums[cur]; l-- {
					}
				}
			}
		}
	}

	return equals
}
