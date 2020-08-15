package Tree

import (
	"strconv"
	"strings"
)

//Given a binary tree, return all root-to-leaf paths.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Input:
//
//1
///   \
//2     3
//\
//5
//
//Output: ["1->2->5", "1->3"]
//
//Explanation: All root-to-leaf paths are: 1->2->5, 1->3

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, path []int, res *[]string) {
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, toString(path))
	}

	if root.Left != nil {
		dfs(root.Left, path, res)
	}

	if root.Right != nil {
		dfs(root.Right, path, res)
	}

}

func toString(path []int) string {
	length := len(path)
	builder := strings.Builder{}
	for i := 0; i < length-1; i++ {
		builder.WriteString(strconv.Itoa(path[i]))
		builder.WriteString("->")
	}
	builder.WriteString(strconv.Itoa(path[length-1]))
	return builder.String()
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	path := []int{}

	dfs(root, path, &res)

	return res
}
