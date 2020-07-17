package _38_RangeSumOfBST

//Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).
//
//The binary search tree is guaranteed to have unique values.
//
// 
//
//Example 1:
//
//Input: root = [10, 5, 15, 3, 7, null, 18], L = 7, R = 15
//Output: 32
//Example 2:
//
//Input: root = [10, 5, 15, 3,7, 13, 18, 1, null,6], L = 6, R = 10
//Output: 23
// 
//
//Note:
//
//The number of nodes in the tree is at most 10000.
//The final answer is guaranteed to be less than 2^31.

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, L int, R int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= L && root.Val <= R {
		*sum += root.Val
	}

	if root.Left != nil {
		dfs(root.Left, L, R, sum)
	}

	if root.Right != nil {
		dfs(root.Right, L, R, sum)
	}

}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int

	dfs(root, L, R, &sum)

	return sum
}
