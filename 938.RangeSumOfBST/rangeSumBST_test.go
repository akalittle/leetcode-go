package _38_RangeSumOfBST

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rangSumBST(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		l    int
		r    int
		want int
	}{
		{
			args: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
			l:    2,
			r:    9,
			want: 16,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, rangeSumBST(item.args, item.l, item.r), "输入:%s", item.args)

	}
}
