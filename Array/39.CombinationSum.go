package Array

//Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target.You may return the combinations in any order.
//
//The same number may be chosen from candidates an unlimited number of times.Two combinations are unique if the frequency of at least one of the chosen numbers is different.
//
//It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
//
// 
//
//Example 1:
//
//Input: candidates = [2, 3, 6, 7], target = 7
//Output: [[2, 2, 3], [7]]
//Explanation:
//2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
//7 is a candidate, and 7 = 7.
//These are the only two combinations.
//Example 2:
//
//Input: candidates = [2, 3, 5], target = 8
//Output: [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
//Example 3:
//
//Input: candidates = [2], target = 1
//Output: []
//Example 4:
//
//Input: candidates = [1], target = 1
//Output: [[1]]
//Example 5:
//
//Input: candidates = [1], target = 2
//Output: [[1, 1]]
// 
//
//Constraints:
//
//1 <= candidates.length <= 30
//1 <= candidates[i] <= 200
//All elements of candidates are distinct.
//1 <= target <= 500

func combinationSum(candidates []int, target int) [][]int {
	comb := []int{}

	ans := [][]int{}

	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			ans = append(ans, tmp)
			return
		}

		// 具体流程可以参考
		//https://assets.leetcode-cn.com/solution-static/39/39_fig1.png

		// 不用这个数
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			// 回溯框架，进行下一步之前加入当前节点，下一步递归过之后，要把当前节点从路径中删除，所以才是“回溯”
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}
