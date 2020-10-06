package LinkedList

//Algorithm of Insertion Sort:
//
//Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
//At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
//It repeats until no input elements remain.
//
//Example 1:
//
//Input: 4->2->1->3
//Output: 1->2->3->4
//Example 2:
//
//Input: -1->5->3->4->0
//Output: -1->0->3->4->5

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head} //假头
	cur := head.Next               //游标
	head.Next = nil                //取消游标的下一个值
	endCur := head
	for cur != nil {
		prev := dummy
		next := cur.Next
		//如果结尾节点数值小于要插入的直接扔到尾巴就好了
		if endCur.Val < cur.Val {
			cur.Next = nil
			endCur.Next = cur
			endCur = cur
			cur = next
			continue
		}
		// 寻找插入点
		for prev.Next != nil && prev.Next.Val <= cur.Val {
			prev = prev.Next
		}
		// prev.Next.Val > p.Val
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	return dummy.Next
}
