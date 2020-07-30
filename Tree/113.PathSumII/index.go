package Tree

//Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given the below binary tree and sum = 22,
//
//5
/// \
//4   8
///   / \
//11  13  4
///  \    / \
//7    2  5   1
//Return:
//
//[
//[5, 4, 11, 2],
//[5, 8, 4, 5]
//]

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum int, path []int, result *[][]int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil && root.Val == sum {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}

	dfs(root.Left, sum-root.Val, path, result)
	dfs(root.Right, sum-root.Val, path, result)

}

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}

	dfs(root, sum, []int{}, &result)

	return result
}
