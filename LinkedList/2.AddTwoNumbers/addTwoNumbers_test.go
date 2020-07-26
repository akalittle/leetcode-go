package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

func Test_AddTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		node1 *ListNode
		node2 *ListNode
		want  *ListNode
	}{
		{
			node1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			node2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, AddTwoNumbers(item.node1, item.node2), "输入:%s", item.node1, item.node2)
	}
}
