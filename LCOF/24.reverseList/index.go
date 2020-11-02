package LCOF

import "github.com/akalittle/leetcode-go/LCOF"

func reverseList(head *LCOF.ListNode) *LCOF.ListNode {

	l := &LCOF.ListNode{}

	current := head

	for current != nil {
		tmp := current.Next
		current.Next = l
		l = current
		current = tmp
	}

	return l
}
