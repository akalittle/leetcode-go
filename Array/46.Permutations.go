package Array

//Given a collection of distinct integers, return all possible permutations.
//
//Example:
//
//Input: [1, 2,3]
//Output:
//[
//[1, 2, 3],
//[1, 3, 2],
//[2, 1, 3],
//[2, 3, 1],
//[3, 1, 2],
//[3, 2, 1]
//]

func permute(nums []int) [][]int {
	var path []int
	var used = make([]bool, len(nums))
	result := [][]int{}
	dfs(nums, path, used, &result)
	return result
}

func dfs(nums []int, path []int, used []bool, result *[][]int) {
	if len(nums) == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == false {
			used[i] = true
			path = append(path, nums[i])

			dfs(nums, path, used, result)

			path = path[:len(path)-1]
			used[i] = false
		}
	}

}
