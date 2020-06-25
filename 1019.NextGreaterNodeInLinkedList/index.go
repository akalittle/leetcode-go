package _1019_NextGreaterNodeInLinkedList

//We are given a linked list with head as the first node.  Let's number the nodes in the list: node_1, node_2, node_3, ... etc.
//
//Each node may have a next larger value: for node_i, next_larger(node_i) is the node_j.val such that j > i, node_j.val > node_i.val, and j is the smallest possible choice.  If such a j does not exist, the next larger value is 0.
//
//Return an array of integers answer, where answer[i] = next_larger(node_{i+1}).
//
//Note that in the example inputs (not outputs) below, arrays such as [2, 1, 5] represent the serialization of a linked list with a head node value of 2, second node value of 1, and third node value of 5.
//
//Example 1:
//
//Input: [2,1,5]
//Output: [5,5,0]
//Example 2:
//
//Input: [2,7,4,3,5]
//Output: [7,0,5,5,0]
//Example 3:
//
//Input: [1,7,5,1,9,2,5,1]
//Output: [7,9,9,9,0,5,0,0]

/**
 * Definition for singly-linked list.
 */

// 单调栈

//从名字上就听的出来，单调栈中存放的数据应该是有序的，所以单调栈也分为单调递增栈和单调递减栈
//
//单调递增栈：栈中数据出栈的序列为单调递增序列
//单调递减栈：栈中数据出栈的序列为单调递减序列
//ps：这里一定要注意所说的递增递减指的是出栈的顺序，而不是在栈中数据的顺序

//现在有一组数10，3，7，4，12。从左到右依次入栈，则如果栈为空或入栈元素值小于栈顶元素值，则入栈；
//否则，如果入栈则会破坏栈的单调性，则需要把比入栈元素小的元素全部出栈。单调递减的栈反之。
//
//10入栈时，栈为空，直接入栈，栈内元素为10。
//
//3入栈时，栈顶元素10比3大，则入栈，栈内元素为10，3。
//
//7入栈时，栈顶元素3比7小，则栈顶元素出栈，此时栈顶元素为10，比7大，则7入栈，栈内元素为10，7。
//
//4入栈时，栈顶元素7比4大，则入栈，栈内元素为10，7，4。
//
//12入栈时，栈顶元素4比12小，4出栈，此时栈顶元素为7，仍比12小，栈顶元素7继续出栈，此时栈顶元素为10，仍比12小，10出栈，此时栈为空，12入栈，栈内元素为12。

type ListNode struct {
	Val  int
	Next *ListNode
}

//使用栈来保存链表数据的索引；
//使栈中的数据保持单调递减。

func NextLargerNodes(head *ListNode) []int {
	var data []int

	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}

	stack, result := make([]int, len(data)), make([]int, len(data))

	for i := 0; i < len(data); i++ {
		if i == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && data[i] > data[stack[len(stack)-1]] {
				pop := len(stack) - 1
				result[stack[pop]] = data[i]
				stack = stack[:pop]
			}
			stack = append(stack, i)
		}
	}

	return result
}
