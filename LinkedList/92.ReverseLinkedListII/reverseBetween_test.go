package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Input: 1->2->3->4->5->NULL, m = 2, n = 4
//Output: 1->4->3->2->5->NULL


func Test_ReverseBetween(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		m    int
		n    int
		want *ListNode
	}{
		{
			args: &ListNode{
				Val: 1,
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
			m: 2,
			n: 4,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, ReverseBetween(item.args, item.m, item.n), "输入:%s", item.args)
	}
}
