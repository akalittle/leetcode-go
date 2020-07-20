package _10_BalancedBinaryTree

//
//Given a binary tree, determine if it is height-balanced.
//
//For this problem, a height-balanced binary tree is defined as:
//
//a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
//
// 
//
//Example 1:
//
//Given the following tree [3, 9, 20, null, null, 15, 7]:
//
//3
/// \
//9  20
///  \
//15   7
//Return true.
//
//Example 2:
//
//Given the following tree [1, 2, 2, 3, 3, null, null, 4, 4]:
//
//1
/// \
//2   2
/// \
//3   3
/// \
//4   4
//Return false.

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	leftH := maxDepth(root.Left)
	rightH := maxDepth(root.Right)
	if abs(leftH-rightH) > 1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
