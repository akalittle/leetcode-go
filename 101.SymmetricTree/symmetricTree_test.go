package _01_SymmetricTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SymmetricTree(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want bool
	}{
		{
			args: &TreeNode{Val: 1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
			},
			want: false,
		},
		{
			args: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: &TreeNode{Val: 2, Left: nil, Right: nil},
			},
			want: true,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, isSymmetric(item.args), "输入:%s", item.args)
		ast.Equal(item.want, isSymmetricWithQueue(item.args), "输入:%s", item.args)

	}
}
