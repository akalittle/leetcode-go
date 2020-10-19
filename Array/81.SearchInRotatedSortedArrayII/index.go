package Array


//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., [0, 0, 1, 2, 2, 5, 6] might become [2, 5, 6, 0, 0,1, 2]).
//
//You are given a target value to search.If found in the array return true, otherwise return false.
//
//Example 1:
//
//Input: nums = [2, 5, 6, 0, 0, 1, 2], target = 0
//Output: true
//Example 2:
//
//Input: nums = [2, 5,6, 0, 0, 1, 2], target = 3
//Output: false
//Follow up:
//
//This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
//Would this affect the run-time complexity? How and why?



func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		} else if nums[l] == nums[mid] {
			l++
			continue
		} else if nums[l] < nums[mid] { // l->mid 有序，比较前半部
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // mid->r有序，比较后半部分
			if nums[r] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
