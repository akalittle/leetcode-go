package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want int
	}{
		{
			args: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
					Right: &TreeNode{Val: 5, Left: nil, Right: nil},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
					Right: nil,
				},
			},
			want: 6,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, countNodes(item.args), "输入:%s", item.args)
	}
}
