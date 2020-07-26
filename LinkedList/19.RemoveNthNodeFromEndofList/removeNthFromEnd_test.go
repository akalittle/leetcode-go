package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//1->2->3->4->5

func Test_RemoveNthFromEnd(t *testing.T) {
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
			val: 2,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
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
	}

	for _, item := range tests {
		ast.Equal(item.want, RemoveNthFromEnd(item.args, item.val), "输入:%s", item.args)
	}
}
