package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		inorder   []int
		postorder []int
		want      *TreeNode
	}{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
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
		ast.Equal(item.want, buildTree(item.inorder, item.postorder), "输入:%s", item.inorder, item.postorder)
	}
}
