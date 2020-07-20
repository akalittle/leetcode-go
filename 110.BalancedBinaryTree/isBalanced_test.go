package _10_BalancedBinaryTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want bool
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
			want: true,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, isBalanced(item.args), "输入:%s", item.args)

	}
}
