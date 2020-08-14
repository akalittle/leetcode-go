package Tree

//Given a complete binary tree, count the number of nodes.
//
//Note:
//
//Definition of a complete binary tree from Wikipedia:
//In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible.It can have between 1 and 2h nodes inclusive at the last level h.
//
//Example:
//
//Input:
//1
/// \
//2   3
/// \  /
//4  5 6
//
//Output: 6

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, count *int) {
	if node == nil {
		return
	}

	*count += 1

	dfs(node.Left, count)
	dfs(node.Right, count)
}

func countNodes(root *TreeNode) int {
	count := 0

	dfs(root, &count)

	return count
}
