package LCOF

import "github.com/akalittle/leetcode-go/LCOF"

//
//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
//例如，二叉树 [1, 2, 2, 3, 4, 4, 3] 是对称的。
//
//1
/// \
//2   2
/// \ / \
//3  4 4  3
//但是下面这个 [1, 2, 2, null, 3, null, 3] 则不是镜像对称的:
//
//1
/// \
//2   2
//\   \
//3    3
//
//
//
//示例 1：
//
//输入：root = [1, 2, 2, 3, 4, 4, 3]
//输出：true
//示例 2：
//
//输入：root = [1,2, 2, null, 3, null,3]
//输出：false

func isSymmetric(root *LCOF.TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)

}

func dfs(a, b *LCOF.TreeNode) bool {

	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && dfs(a.Left, b.Right) && dfs(a.Right, b.Left)
}
