package Tree

/**
 * Definition for a binary tree node.
 */

//Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

//					For example:
//					Given binary tree [3, 9, 20, null, null, 15, 7],
//					 3
//					/ \
//					9  20
//				    /  \
//					15   7
//					return its level order traversal as:
//					[
//					[3],
//					[9, 20],
//					[15, 7]
//					]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res

}
