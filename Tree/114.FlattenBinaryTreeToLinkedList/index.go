package Tree

//
//Given a binary tree, flatten it to a linked list in-place.
//
//For example, given the following tree:
//
//1
/// \
//2   5
/// \   \
//3   4   6
//The flattened tree should look like:
//
//1
//\
//2
//\
//3
//\
//4
//\
//5
//\
//6

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {

	if root == nil {
		return nil
	}

	r1 := preorderTraversal(root.Left)
	r2 := preorderTraversal(root.Right)

	return append(append([]*TreeNode{root}, r1...), r2...)

}
