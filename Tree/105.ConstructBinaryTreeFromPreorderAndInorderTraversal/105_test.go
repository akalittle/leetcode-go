package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		preorder []int

		inorder []int
		want    *TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},

			inorder: []int{9, 3, 15, 20, 7},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, buildTree(item.preorder, item.inorder), "输入:%s", item.inorder, item.preorder)
	}
}
