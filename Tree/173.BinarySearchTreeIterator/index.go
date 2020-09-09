package Tree

//Implement an iterator over a binary search tree (BST).Your iterator will be initialized with the root node of a BST.
//
//Calling next() will return the next smallest number in the BST.
//
//
//Example:
//
//BSTIterator iterator = new BSTIterator(root)
//iterator.next()    // return 3
//iterator.next()    // return 7
//iterator.hasNext() // return true
//iterator.next()    // return 9
//iterator.hasNext() // return true
//iterator.next()    // return 15
//iterator.hasNext() // return true
//iterator.next()    // return 20
//iterator.hasNext() // return false
// 
//
//Note:
//
//next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
//You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root *TreeNode
	nums []int
}

func Constructor(root *TreeNode) BSTIterator {
	nums := make([]int, 0)
	inOrder(root, &nums)

	return BSTIterator{
		root: root,
		nums: nums,
	}

}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	smallest := this.nums[0]
	// 删除第一个元素
	this.nums = this.nums[1:]
	return smallest
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.nums) > 0 {
		return true
	}
	return false
}

func inOrder(root *TreeNode, n *[]int) {
	if root == nil {
		return
	}
	// 左
	inOrder(root.Left, n)
	// 中
	*n = append(*n, root.Val)
	// 右
	inOrder(root.Right, n)

}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
