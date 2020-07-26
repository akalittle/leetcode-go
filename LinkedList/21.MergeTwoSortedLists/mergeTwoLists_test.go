package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddTwoNumbers(t *testing.T) {
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
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			node2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, MergeTwoLists(item.node1, item.node2), "输入:%s", item.node1, item.node2)
	}
}
