package LCOF

import (
	"github.com/akalittle/leetcode-go/LCOF"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mirrorTree(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *LCOF.TreeNode
		want *LCOF.TreeNode
	}{
		{
			args: &LCOF.TreeNode{
				Val: 4,
				Left: &LCOF.TreeNode{
					Val: 2,
					Left: &LCOF.TreeNode{
						Val: 1,
					},
					Right: &LCOF.TreeNode{
						Val: 3,
					},
				},
				Right: &LCOF.TreeNode{
					Val: 7,
					Left: &LCOF.TreeNode{
						Val: 6,
					},
					Right: &LCOF.TreeNode{
						Val: 9,
					},
				},
			},
			want: &LCOF.TreeNode{
				Val: 4,
				Left: &LCOF.TreeNode{
					Val: 7,
					Left: &LCOF.TreeNode{
						Val: 9,
					},
					Right: &LCOF.TreeNode{
						Val: 6,
					},
				},
				Right: &LCOF.TreeNode{
					Val: 2,
					Left: &LCOF.TreeNode{
						Val: 3,
					},
					Right: &LCOF.TreeNode{
						Val: 1,
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, mirrorTree(item.args), "输入:%s", item.args)
	}

}
