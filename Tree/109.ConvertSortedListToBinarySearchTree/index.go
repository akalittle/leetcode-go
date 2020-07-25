package Tree

//Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//Example:
//
//Given the sorted linked list: [-10, -3, 0, 5, 9],
//
//One possible answer is: [0, -3, 9, -10, null, 5], which represents the following height balanced BST:
//
//0
/// \
//-3   9
///   /
//-10  5

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := findMiddleNode(head)
	node := &TreeNode{
		Val: mid.Val,
	}

	if head == mid {
		return node
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)
	return node

}

func findMiddleNode(l *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := l, l
	// 快慢指针找 链表的中间节点
	for fast != nil && fast.Next != nil { //end or end.Next
		prev = slow
		slow = slow.Next      //move 1
		fast = fast.Next.Next //move 2
	}

	if prev != nil {
		prev.Next = nil
	}
	return slow

}

func sortedListToBST(head *ListNode) *TreeNode {
	return helper(head)
}
