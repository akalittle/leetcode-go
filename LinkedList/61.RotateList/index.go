package LinkedList

//
//Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
//Example 1:
//
//Input: 1->2->3->4->5->NULL, k = 2
//Output: 4->5->1->2->3->NULL
//Explanation:
//rotate 1 steps to the right: 5->1->2->3->4->NULL
//rotate 2 steps to the right: 4->5->1->2->3->NULL
//Example 2:
//
//Input: 0->1->2->NULL, k = 4
//Output: 2->0->1->NULL
//Explanation:
//rotate 1 steps to the right: 2->0->1->NULL
//rotate 2 steps to the right: 1->2->0->NULL
//rotate 3 steps to the right: 0->1->2->NULL
//rotate 4 steps to the right: 2->0->1->NULL

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 循环找到链表长度n和链表尾，然后算一下k/n的余数，
// 这时候把链表连成环，然后从尾部开始往前走n-k步，在这里断开这个环，然后返回头指针。

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	n, p := 1, head
	for p.Next != nil {
		p = p.Next
		n++
	}
	p.Next = head
	k %= n
	for i := 1; i <= n-k; i++ {
		p = p.Next
	}
	head, p.Next = p.Next, nil
	return head
}
