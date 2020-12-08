package LCOF

import (
	"github.com/akalittle/leetcode-go/LCOF"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *LCOF.ListNode
		want []int
	}{
		{
			args: &LCOF.ListNode{
				Val: 1,
				Next: &LCOF.ListNode{
					Val: 3,
					Next: &LCOF.ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: []int{2, 3, 1},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, reversePrint(item.args), "输入:%s", item.args)
	}
}
