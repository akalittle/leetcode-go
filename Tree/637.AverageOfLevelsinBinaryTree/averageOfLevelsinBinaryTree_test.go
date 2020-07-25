package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AverageOfLevelsinBinaryTree(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want []float64
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
			want: []float64{
				3, 14.5, 11,
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, averageOfLevels(item.args), "输入:%s", item.args)

	}
}
