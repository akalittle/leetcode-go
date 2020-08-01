package Array

//Given two arrays, write a function to compute their intersection.
//
//Example 1:
//
//Input: nums1 = [1, 2, 2, 1], nums2 = [2, 2]
//Output: [2]
//Example 2:
//
//Input: nums1 = [4, 9, 5], nums2 = [9, 4, 9, 8, 4]
//Output: [9, 4]
//Note:
//
//Each element in the result must be unique.
//The result can be in any order.

func intersection(nums1 []int, nums2 []int) []int {
	var res = []int{}

	hash := make(map[int]bool)

	for _, v := range nums1 {
		hash[v] = true
	}

	for _, v := range nums2 {
		if hash[v] {
			res = append(res, v)
			hash[v] = false
		}
	}

	return res
}
