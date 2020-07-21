package _08_ConvertSortedArrayToBinarySearchTree

//Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//Example:
//
//Given the sorted array: [-10, -3, 0, 5,9],
//
//One possible answer is: [0, -3, 9, -10, null, 5], which represents the following height balanced BST:
//
//0
/// \
//-3   9
///   /
//-10  5

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2

	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root

}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}
