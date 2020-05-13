package _21_MergeTwoSortedLists_

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	guard := &ListNode{}

	result := guard

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			guard.Next = l1
			l1 = l1.Next
		} else {
			guard.Next = l2
			l2 = l2.Next
		}
		guard = guard.Next
	}

	if l1 != nil {
		guard.Next = l1
	}

	if l2 != nil {
		guard.Next = l2
	}

	return result.Next
}
