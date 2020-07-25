package LinkedList

//Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
//
//You should preserve the original relative order of the nodes in each of the two partitions.
//
//Example:
//
//Input: head = 1->4->3->2->5->2, x = 3
//Output: 1->2->2->4->3->5

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func Partition(head *ListNode, x int) *ListNode {
	// 创建两个链表
	// greaterDummy 存放链表中较大的数据
	// lessDummy 存放联保中较小的数据
	// 最终把两个链表拼接起来就好
	greaterDummy := &ListNode{}
	lessDummy := &ListNode{}
	curGreater := greaterDummy
	curLess := lessDummy

	for head != nil {
		if head.Val < x {
			curLess.Next = head
			curLess = head
		} else {
			curGreater.Next = head
			curGreater = head
		}
		head = head.Next
	}

	curGreater.Next = nil
	curLess.Next = greaterDummy.Next
	return lessDummy.Next
}
