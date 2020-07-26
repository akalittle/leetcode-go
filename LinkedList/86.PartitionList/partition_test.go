package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Input: head = 1->4->3->2->5->2, x = 3
//Output: 1->2->2->4->3->5

func Test_Partition(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		val  int
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
			val: 3,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
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
		ast.Equal(item.want, Partition(item.args, item.val), "输入:%s", item.args)
	}
}
