package _206_ReverseLinkedList

//Reverse a singly linked list.
//
//Example:
//
//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL

/**
* Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newList *ListNode
	current := head

	for current != nil {
		temp := current.Next
		current.Next = newList
		newList = current
		current = temp
	}

	return newList
}
