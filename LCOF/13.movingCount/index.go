package LCOF

//地上有一个m行n列的方格，从坐标 [0, 0] 到坐标 [m-1, n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7 = 18。但它不能进入方格 [35, 38]，因为3+5+3+8 = 19。请问该机器人能够到达多少个格子？
//
// 
//
//示例 1：
//
//输入：m = 2, n = 3, k = 1
//输出：3
//示例 2：
//
//输入：m = 3, n = 1, k = 0
//输出：1

func sum(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}

func inBound(m, n, i, j, k int, visited [][]bool) bool {
	return i >= 0 && i < m && j >= 0 && j < n && sum(i)+sum(j) <= k && !visited[i][j]
}

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return dfs(m, n, 0, 0, k, visited)
}

//
func dfs(m, n, i, j, k int, visited [][]bool) int {
	if !inBound(m, n, i, j, k, visited) {
		return 0
	}

	visited[i][j] = true

	sum := 1

	sum += dfs(m, n, i+1, j, k, visited)
	sum += dfs(m, n, i-1, j, k, visited)
	sum += dfs(m, n, i, j+1, k, visited)
	sum += dfs(m, n, i, j-1, k, visited)

	return sum
}
