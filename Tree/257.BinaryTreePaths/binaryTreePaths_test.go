package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want []string
	}{
		{
			args: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: []string{"1->2->5", "1->3"},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, binaryTreePaths(item.args), "输入:%s", item.args)
	}
}
