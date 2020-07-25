package Tree

//Given a binary tree, return the sum of values of nodes with even-valued grandparent.  (A grandparent of a node is the parent of its parent, if it exists.)
//
//If there are no nodes with an even-valued grandparent, return 0.
//
// 
//
//Example 1:
//
//
//
//Input: root = [6, 7, 8, 2,7, 1, 3, 9, null,1, 4, null, null, null,5]
//Output: 18
//Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, parent *TreeNode, grand *TreeNode, sum *int) {
	if root != nil {
		if grand != nil && grand.Val%2 == 0 {
			*sum = *sum + root.Val
		}

		dfs(root.Left, root, parent, sum)
		dfs(root.Right, root, parent, sum)
	}

}

func sumEvenGrandparent(root *TreeNode) int {

	sum := 0

	dfs(root, nil, nil, &sum)

	return sum
}
