package _94_BinaryTreeInorderTraversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinaryTreeInorderTraversal(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want []int
	}{
		{
			args: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}}},
			want: []int{1, 3, 2},
		},
		{
			args: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			want: []int{2, 1, 3},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, InorderTraversal(item.args), "输入:%s", item.args)

		ast.Equal(item.want, InOrderTraversalWithStack(item.args), "输入:%s", item.args)
	}
}
