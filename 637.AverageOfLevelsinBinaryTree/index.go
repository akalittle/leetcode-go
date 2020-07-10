package _37_AverageOfLevelsinBinaryTree

/**
 * Definition for a binary tree node.

 */

// 其实就是层序遍历 求平均值

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := [][]int{}
	if root == nil {
		return []float64{}
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

	var avg []float64

	for _, v := range res {
		var total float64
		for _, item := range v {
			total += float64(item)
		}
		avg = append(avg, total/float64(len(v)))
	}

	return avg
}
