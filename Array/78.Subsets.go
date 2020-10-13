package Array

//Given a set of distinct integers, nums, return all possible subsets (the power set).
//
//Note: The solution set must not contain duplicate subsets.
//
//Example:
//
//Input: nums = [1, 2,3]
//Output:
//[
//[3],
//  [1],
//  [2],
//  [1, 2, 3],
//  [1, 3],
//  [2,3],
//  [1, 2],
//  []
//]

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	var backtrace func(pos, num int, cur []int)

	backtrace = func(p int, num int, cur []int) {
		if len(cur) == num {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		for i := p; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrace(i+1, num, cur)
			cur = cur[:len(cur)-1]
		}

	}

	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		backtrace(0, i, cur)
	}

	return ans
}
