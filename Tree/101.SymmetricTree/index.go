package Tree

//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
//For example, this binary tree [1,2, 2, 3, 4, 4,3] is symmetric:
//
//1
/// \
//2   2
/// \ / \
//3  4 4  3
// 
//
//But the following [1, 2,2, null, 3, null, 3] is not:
//
//1
/// \
//2   2
//\   \
//3    3
// 
//
//Follow up: Solve it both recursively and iteratively.

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func isSymmetricWithQueue(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
