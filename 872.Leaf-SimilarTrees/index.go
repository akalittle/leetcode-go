package _72_Leaf_SimilarTrees

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, leaves *[]int) {

	if root != nil {
		dfs(root.Left, leaves)
		if root.Left == nil && root.Right == nil {
			*leaves = append(*leaves, root.Val)
		}
		dfs(root.Right, leaves)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	t1 := []int{}

	t2 := []int{}

	dfs(root1, &t1)

	dfs(root2, &t2)

	if len(t1) != len(t2) {
		return false
	}
	for i, _ := range t1 {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true

}
