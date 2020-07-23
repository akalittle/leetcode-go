package _11_MinimumDepthOfBinaryTree

//Given a binary tree, find its minimum depth.
//
//The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given binary tree [3, 9, 20, null, null,15, 7],
//
//3
/// \
//9  20
///  \
//15   7

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	if node.Left != nil && node.Right == nil {
		return dfs(node.Left) + 1
	}

	if node.Left == nil && node.Right != nil {
		return dfs(node.Right) + 1
	}

	leftHeight := dfs(node.Left)
	rightHeight := dfs(node.Right)

	return min(leftHeight, rightHeight) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	return dfs(root)
}
