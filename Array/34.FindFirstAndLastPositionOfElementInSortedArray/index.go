package Array

//Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
//
//Your algorithm's runtime complexity must be in the order of O(log n).
//
//If the target is not found in the array, return [-1, -1].
//
//Example 1:
//
//Input: nums = [5, 7, 7, 8, 8, 10], target = 8
//Output: [3, 4]
//Example 2:
//
//Input: nums = [5, 7,7, 8, 8, 10], target = 6
//Output: [-1, -1]
// 
//
//Constraints:
//
//0 <= nums.length <= 10^5
//-10^9 <= nums[i] <= 10^9
//nums is a non decreasing array.
//-10^9 <= target <= 10^9

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	start, end := -1, -1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {

			start, end = mid, mid
			// 在mid的前后范围内进行查询
			for start > 0 && nums[start] == nums[start-1] {
				start--
			}
			for end < len(nums)-1 && nums[end] == nums[end+1] {
				end++
			}
			break

		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1

		}
	}

	return []int{start, end}

}
