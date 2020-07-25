package LinkedList

//Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
//
//Return the linked list sorted as well.
//
//Example 1:
//
//Input: 1->2->3->3->4->4->5
//Output: 1->2->5
//Example 2:
//
//Input: 1->1->1->2->3
//Output: 2->3

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	prev := dummy

	cur := dummy.Next

	for cur != nil && cur.Next != nil {

		if cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}
