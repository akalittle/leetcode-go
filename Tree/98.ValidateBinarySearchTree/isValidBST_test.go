package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinaryTreeInorderTraversal(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want bool
	}{
		{
			args: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: false,
		},
		{
			args: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: true,
		},
		{
			args: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   17,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: false,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, isValidBST(item.args), "输入:%s", item.args)
	}
}
