package Tree

//
//Given inorder and postorder traversal of a tree, construct the binary tree.
//
//Note:
//You may assume that duplicates do not exist in the tree.
//
//For example, given
//
//inorder =  [9, 3,15, 20, 7]
//postorder = [9, 15, 7, 20, 3]
//Return the following binary tree:
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	if len(inorder) == 1 || len(postorder) == 1 {
		return root
	}

	for i, val := range inorder {
		if val == postorder[len(postorder)-1] {
			root.Left = buildTree(inorder[:i], postorder[:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
		}
	}

	return root
}
