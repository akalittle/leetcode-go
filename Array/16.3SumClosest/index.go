package Array

import (
	"math"
	"sort"
)

//Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.Return the sum of the three integers.You may assume that each input would have exactly one solution.
//
// 
//
//Example 1:
//
//Input: nums = [-1, 2, 1, -4], target = 1
//Output: 2
//Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
// 
//
//Constraints:
//
//3 <= nums.length <= 10^3
//-10^3 <= nums[i] <= 10^3
//-10^4 <= target <= 10^4

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	n := len(nums)

	best := math.MaxInt32

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		second := first + 1

		third := n - 1

		for second < third {
			sum := nums[first] + nums[second] + nums[third]

			if sum == target {
				return target
			}

			best = update(sum, best, target)

			if sum > target {
				// 如果和大于 target，移动 third 对应的指针
				third2 := third - 1
				// 移动到下一个不相等的元素
				for second < third2 && nums[third2] == nums[third] {
					third2--
				}
				third = third2
			} else {
				// 如果和小于 target，移动 second 对应的指针
				second2 := second + 1
				// 移动到下一个不相等的元素
				for second2 < third && nums[second2] == nums[second] {
					second2++
				}
				second = second2
			}

		}

	}

	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func update(cur int, best int, target int) (be int) {

	if abs(cur-target) < abs(best-target) {
		best = cur
	}

	return best
}
