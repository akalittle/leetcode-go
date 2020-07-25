package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		tree1 *TreeNode
		tree2 *TreeNode
		want  *TreeNode
	}{
		{
			tree1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &TreeNode{Val: 2, Left: nil, Right: nil},
			},
			tree2: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
					Right: &TreeNode{Val: 4, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, mergeTrees(item.tree1, item.tree2), "")
	}
}
