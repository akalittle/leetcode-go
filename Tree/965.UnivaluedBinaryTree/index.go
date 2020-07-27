package Tree

//A binary tree is univalued if every node in the tree has the same value.
//
//Return true if and only if the given tree is univalued.

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, val int) bool {

	if root != nil {

		if root.Val != val {
			return false
		}
		return dfs(root.Left, val) && dfs(root.Right, val)

	}
	return true

}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return dfs(root, root.Val)
}
