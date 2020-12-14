package LCOF

import "github.com/akalittle/leetcode-go/LCOF"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *LCOF.ListNode, l2 *LCOF.ListNode) *LCOF.ListNode {
	guard := &LCOF.ListNode{}

	result := guard

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			guard.Next = l2
			l2 = l2.Next
		} else {
			guard.Next = l1
			l1 = l1.Next
		}
	}

	if l1 != nil {
		guard.Next = l1
	}
	if l2 != nil {
		guard.Next = l2
	}

	return result.Next
}
