package LinkedList

//Given a linked list, swap every two adjacent nodes and return its head.
//
//You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// 
//
//Example:
//
//Given 1->2->3->4, you should return the list as 2->1->4->3.

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// 第一个节点和第二个节点为空时，直接返回原链表
	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	//交换firstNode与second,
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}
