package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BSTIterator(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
	}{
		{
			args: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   20,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	for _, item := range tests {

		iterator := Constructor(item.args)

		ast.Equal(3, iterator.Next())
		ast.Equal(7, iterator.Next())
		ast.Equal(true, iterator.HasNext())
		ast.Equal(9, iterator.Next())
		ast.Equal(true, iterator.HasNext())
		ast.Equal(15, iterator.Next())
		ast.Equal(20, iterator.Next())
		ast.Equal(false, iterator.HasNext())
	}
}
