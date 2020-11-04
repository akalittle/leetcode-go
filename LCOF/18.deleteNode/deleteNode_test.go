package LCOF

import (
	"github.com/akalittle/leetcode-go/LCOF"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *LCOF.ListNode
		val  int
		want *LCOF.ListNode
	}{
		{
			args: &LCOF.ListNode{
				Val: 4,
				Next: &LCOF.ListNode{
					Val: 5,
					Next: &LCOF.ListNode{
						Val: 1,
						Next: &LCOF.ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
			val: 5,
			want: &LCOF.ListNode{
				Val: 4,
				Next: &LCOF.ListNode{
					Val: 1,
					Next: &LCOF.ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, deleteNode(item.args, item.val), "输入:%s", item.args)
	}
}
