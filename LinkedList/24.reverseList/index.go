package LCOF

import "github.com/akalittle/leetcode-go/LCOF"

//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
// 
//
//限制：
//
//0 <= 节点个数 <= 5000

func reverseList(head *LCOF.ListNode) *LCOF.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	d := &LCOF.ListNode{}

	for head != nil {
		tmp := head.Next
		head.Next = d
		d = head
		head = tmp
	}

	return d
}
