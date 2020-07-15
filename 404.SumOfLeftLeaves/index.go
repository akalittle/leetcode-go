package _04_SumOfLeftLeaves

//Find the sum of all left leaves in a given binary tree.
//
//Example:
//
//3
/// \
//9  20
///  \
//15   7
//
//There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}

	sum := 0

	if root.Left != nil {
		// 找到左侧叶子
		if root.Left.Left == nil && root.Left.Right == nil {
			sum = root.Left.Val
		}

		sum += sumOfLeftLeaves(root.Left)
	}

	if root.Right != nil {
		sum += sumOfLeftLeaves(root.Right)
	}

	return sum
}
