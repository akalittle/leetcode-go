package _72_Leaf_SimilarTrees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		tree1 *TreeNode
		tree2 *TreeNode

		want bool
	}{
		{
			tree1: &TreeNode{
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
			tree2: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			},
			want: false,
		},
		{
			tree1: &TreeNode{
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
			tree2: &TreeNode{
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
		ast.Equal(item.want, leafSimilar(item.tree1, item.tree2), "输入:%s", item.tree1, item.tree2)
	}
}
