package LCOF

import (
	"github.com/akalittle/leetcode-go/LCOF"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *LCOF.ListNode
		k    int
		want *LCOF.ListNode
	}{
		{
			args: &LCOF.ListNode{
				Val: 1,
				Next: &LCOF.ListNode{
					Val: 2,
					Next: &LCOF.ListNode{
						Val: 3,
						Next: &LCOF.ListNode{
							Val: 4,
							Next: &LCOF.ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			k: 2,
			want: &LCOF.ListNode{
				Val: 4,
				Next: &LCOF.ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, getKthFromEnd(item.args, item.k), "输入:%s", item.args)
	}
}
