package _144_BinaryTreePreorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func helper(node *TreeNode) int[] {
//
//return int[]{1, 2}
//}

func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r1 := PreOrderTraversal(root.Left)
	r2 := PreOrderTraversal(root.Right)
	return append(append([]int{root.Val}, r1...), r2...)
}

func PreOrderTraversalWithStack(root *TreeNode) []int {

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

		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}
	return res

}
