package LinkedList

//Given a sorted linked list, delete all duplicates such that each element appear only once.
//
//Example 1:
//
//Input: 1->1->2
//Output: 1->2
//Example 2:
//
//Input: 1->1->2->3->3
//Output: 1->2->3

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head.Next

	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}

	slow.Next = nil
	return head
}
