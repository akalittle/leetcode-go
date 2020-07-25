package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchBST(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want *TreeNode
	}{
		{
			args: []int{
				1, 2, 3,
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, sortedArrayToBST(item.args), "输入:%s", item.args)

	}
}
