package Tree

//Given a binary tree, return the inorder traversal of its nodes' values.
//
//Example:
//
//Input: [1, null, 2, 3]
// 1
//  \
//   2
//  /
// 3
//
//Output: [1, 3, 2]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func helper(node *TreeNode) int[] {
//
//return int[]{1, 2}
//}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r1 := InorderTraversal(root.Left)
	r2 := InorderTraversal(root.Right)
	return append(append(r1, root.Val), r2...)
}

func InOrderTraversalWithStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for len(stack) > 0 || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}
		index := len(stack) - 1             //栈顶
		res = append(res, stack[index].Val) //中序输出
		root = stack[index].Right           //右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:index]               //出栈
	}
	return res

}
