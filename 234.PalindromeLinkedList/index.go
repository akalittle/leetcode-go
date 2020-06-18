package _234_PalindromeLinkedList

//避免使用 O(n)O(n) 额外空间的方法就是改变输入。
//
//我们可以将链表的后半部分反转（修改链表结构），然后将前半部分和后半部分进行比较。比较完成后我们应该将链表恢复原样。虽然不需要恢复也能通过测试用例，因为使用该函数的人不希望链表结构被更改。
//
//算法：
//
//我们可以分为以下几个步骤：
//
//找到前半部分链表的尾节点。
//反转后半部分链表。
//判断是否为回文。
//恢复链表。
//返回结果。
//执行步骤一，我们可以计算链表节点的数量，然后遍历链表找到前半部分的尾节点。
//
//或者可以使用快慢指针在一次遍历中找到：慢指针一次走一步，快指针一次走两步，快慢指针同时出发。当快指针移动到链表的末尾时，慢指针到链表的中间。通过慢指针将链表分为两部分。
//
//若链表有奇数个节点，则中间的节点应该看作是前半部分。
//
//步骤二可以使用在反向链表问题中找到解决方法来反转链表的后半部分。
//
//步骤三比较两个部分的值，当后半部分到达末尾则比较完成，可以忽略计数情况中的中间节点。
//
//步骤四与步骤二使用的函数相同，再反转一次恢复链表本身。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	current := head
	for current != nil {
		tmp := current.Next
		current.Next = pre
		pre = current
		current = tmp
	}
	return pre

}

func IsPalindrome(head *ListNode) bool {
	// 定义快慢指针
	slow := head
	fast := head
	if head == nil || head.Next == nil {
		return true
	}

	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversedRightList := reverseLinkedList(slow.Next)

	p1 := head
	p2 := reversedRightList

	//从短的链表开始比较
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true

}
