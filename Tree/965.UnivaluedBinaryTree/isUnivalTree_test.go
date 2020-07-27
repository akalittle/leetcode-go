package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rangSumBST(t *testing.T) {
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
			want: false,
		},
		{
			args: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			want: true,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, isUnivalTree(item.args), "输入:%s", item.args)
	}
}
