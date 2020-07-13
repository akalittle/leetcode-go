package _26_InvertBinaryTree

//Example:
//
//Input:
//
//     4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
//Output:
//
//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	right := invertTree(root.Right)
	left := invertTree(root.Left)

	root.Left = right
	root.Right = left

	return root
}
