package _92_ReverseLinkedListII_

//Reverse a linked list from position m to n. Do it in one-pass.
//
//Note: 1 ≤ m ≤ n ≤ length of list.
//
//Example:
//
//Input: 1->2->3->4->5->NULL, m = 2, n = 4
//Output: 1->4->3->2->5->NULL

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {

	if head == nil || m > n {
		return nil
	}

	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	//走到将要翻转节点的前一个节点 prev
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}

	//cur 第m个节点，也就是将要翻转的节点
	cur := prev.Next
	for i := m; i < n; i++ {
		tmp := cur.Next      //保存要反转节点的下一个节点
		cur.Next = tmp.Next  //当前节点指向 要放转节点的next节点，最终指向第m个节点的next
		tmp.Next = prev.Next //第n个节点的next指向前一个节点
		prev.Next = tmp      // 第m个节点指向后面一个节点
	}
	return dummy.Next
}
