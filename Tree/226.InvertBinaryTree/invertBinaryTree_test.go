package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InvertBinaryTree(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want *TreeNode
	}{
		{
			args: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: nil,
				},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{Val: 3, Left: nil, Right: nil},
				},
				Right: nil,
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, invertTree(item.args), "输入:%s", item.args)

	}
}
