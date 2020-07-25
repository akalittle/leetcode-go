package Array

//A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.
//
//Now given an M x N matrix, return  True if and only if the matrix is Toeplitz.
// 
//
//Example 1:
//
//Input:
//matrix = [
//  [1, 2, 3, 4],
//  [5, 1, 2, 3],
//  [9, 5, 1, 2]
//]
//Output: True
//Explanation:
//In the above grid, the diagonals are:
//"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
//In each diagonal all elements are the same, so the answer is True.
//Example 2:
//
//Input:
//matrix = [
//  [1, 2],
//  [2, 2]
//]
//Output: False
//Explanation:
//The diagonal "[1, 2]" has different elements.

func IsToeplitzMatrix(matrix [][]int) bool {
	hash := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if _, ok := hash[i-j]; ok {
				if hash[i-j] != matrix[i][j] {
					return false
				}
			} else {
				hash[i-j] = matrix[i][j]
			}
		}
	}
	return true

}
