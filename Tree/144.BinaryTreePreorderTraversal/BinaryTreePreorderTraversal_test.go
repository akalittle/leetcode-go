package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want []int
	}{
		{
			args: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}}},
			want: []int{1, 2, 3},
		},
		{
			args: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			want: []int{1, 2, 3},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, PreOrderTraversal(item.args), "输入:%s", item.args)
		ast.Equal(item.want, PreOrderTraversalWithStack(item.args), "输入:%s", item.args)
	}
}
