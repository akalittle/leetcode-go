package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Input: head = 1->4->3->2->5->2, x = 3
//Output: 1->2->2->4->3->5

// todo

func Test_GetIntersectionNode(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		node1 *ListNode
		node2 *ListNode
		want  *ListNode
	}{
		{
			node1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			node2: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, GetIntersectionNode(item.node1, item.node2), "输入:%s", item.node1, item.node2)
	}
}
