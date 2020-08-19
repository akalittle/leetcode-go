package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_flatten(t *testing.T) {
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
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, item := range tests {

		flatten(item.args)
		ast.Equal(item.want, item.args)
	}
}
