package LinkedList

//Given a singly linked list L: L0→L1→…→Ln-1→Ln,
//reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//You may not modify the values in the list's nodes, only nodes itself may be changed.
//
//Example 1:
//
//Given 1->2->3->4, reorder it to 1->4->2->3.
//Example 2:
//
//Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

/**
 * Definition for singly-linked list.
 */

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	r := reverseLinkedList(slow.Next)

	slow.Next = nil

	//&fmt

	//fmt.Println("head", head)

	for r != nil {
		temp := r.Next
		r.Next = head.Next
		head.Next = r
		head = r.Next
		r = temp
	}

}

func reverseLinkedList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	current := head

	for current != nil {
		tmp := current.Next
		current.Next = pre
		pre = current
		current = tmp
	}
	return pre

}
