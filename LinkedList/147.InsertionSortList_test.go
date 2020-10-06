package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MiddleNode(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		want *ListNode
	}{
		{
			args: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  2,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, insertionSortList(item.args), "输入:%s", item.args)
	}
}
