package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want [][]int
	}{
		{
			args: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9, Left: nil, Right: nil},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
					Right: &TreeNode{Val: 7, Left: nil, Right: nil},
				}},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, levelOrder(item.args), "输入:%s", item.args)

	}
}
