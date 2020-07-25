package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchBST(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		val  int
		want *TreeNode
	}{
		{
			args: &TreeNode{
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
			val: 2,
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, searchBST(item.args, item.val), "输入:%s", item.args)
		ast.Equal(item.want, searchBSTIterator(item.args, item.val), "输入:%s", item.args)

	}
}
