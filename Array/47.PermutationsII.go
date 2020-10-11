package Array

import "sort"

//Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//
//Example:
//
//Input: [1, 1, 2]
//Output:
//[
//[1, 1, 2],
//[1, 2, 1],
//[2, 1, 1]
//]

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			tmp := make([]int, len(perm))
			copy(tmp, perm)
			ans = append(ans, tmp)
			return
		}
		for i, v := range nums {
			//
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
