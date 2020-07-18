package _315_SumOfNodesWithEven_ValuedGrandparent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumEvenGrandparent(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		want int
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
			want: 4,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, sumEvenGrandparent(item.args), "输入:%s", item.args)
	}
}
