package _145_BinaryTreePostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r1 := PostOrderTraversal(root.Left)
	r2 := PostOrderTraversal(root.Right)
	return append(append(r1, r2...), root.Val)
}

func PostOrderTraversalWithStack(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var res []int
	var stack []*TreeNode

	stack = append(stack, root)

	for len(stack) > 0 {

		index := len(stack) - 1 //栈顶
		node := stack[index]
		stack = stack[:index] //出栈

		res = append([]int{node.Val}, res...)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}
	return res

}
