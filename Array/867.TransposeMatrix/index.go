package Array

//
//Given a 2D integer array matrix, return the transpose of matrix.
//
//The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.
//
//
//Example 1:
//
//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [[1,4,7],[2,5,8],[3,6,9]]
//Example 2:
//
//Input: matrix = [[1,2,3],[4,5,6]]
//Output: [[1,4],[2,5],[3,6]]

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])

	t := make([][]int, m)

	for i := range t {
		t[i] = make([]int, n)
		for j := range t[i] {
			t[i][j] = -1
		}
	}

	for i, row := range matrix {
		for j, v := range row {
			t[j][i] = v
		}
	}
	return t
}
