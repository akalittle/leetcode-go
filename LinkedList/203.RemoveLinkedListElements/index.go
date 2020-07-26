package LinkedList

//Remove all elements from a linked list of integers that have value val.
//
//Example:
//
//Input:  1->2->6->3->4->5->6, val = 6
//Output: 1->2->3->4->5

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {

	p := head

	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	if head != nil && head.Val == val {
		p = p.Next
	}

	return head
}

func RemoveElementsWithDummyHead(head *ListNode, val int) *ListNode {
	node := new(ListNode)
	node.Next = head
	pre, cur := node, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return node.Next
}

func RemoveElementsWithRecursive(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = RemoveElementsWithRecursive(head.Next, val)

	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
