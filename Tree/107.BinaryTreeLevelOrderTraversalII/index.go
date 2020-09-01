package Tree

//Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
//
//For example:
//Given binary tree [3, 9, 20,null, null, 15, 7],
//3
/// \
//9  20
///  \
//15   7
//return its bottom-up level order traversal as:
//[
//[15, 7],
//[9, 20],
//[3]
//]

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}

	dfs(root, &res, 0)

	l := len(res)
	end := l / 2
	for i := 0; i < end; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}

	return res

}

func dfs(root *TreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)

	dfs(root.Left, res, level+1)
	dfs(root.Right, res, level+1)

}
