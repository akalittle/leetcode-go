package Tree

//Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
//
//For example:
//Given binary tree [3, 9, 20, null,null, 15, 7],
//3
/// \
//9  20
///  \
//15   7
//return its zigzag level order traversal as:
//[
//[3],
//[20, 9],
//[15, 7]
//]

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, r *[][]int, layer int) {
	if root == nil {
		return
	}

	if len(*r) <= layer {
		*r = append(*r, []int{})
	}

	// 偶数
	if layer%2 == 0 {
		(*r)[layer] = append((*r)[layer], root.Val)
	} else {
		(*r)[layer] = append([]int{root.Val}, (*r)[layer]...)
	}
	dfs(root.Left, r, layer+1)
	dfs(root.Right, r, layer+1)

}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	dfs(root, &result, 0)

	return result
}
