package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			args: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &TreeNode{Val: 5, Left: nil, Right: nil},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
					Right: &TreeNode{Val: 9, Left: nil, Right: nil},
				},
			},
			p: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0, Left: nil, Right: nil},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
					Right: &TreeNode{Val: 5, Left: nil, Right: nil},
				},
			},
			q: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
				Right: &TreeNode{Val: 9, Left: nil, Right: nil},
			},
			want: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0, Left: nil, Right: nil},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
						Right: &TreeNode{Val: 5, Left: nil, Right: nil},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
					Right: &TreeNode{Val: 9, Left: nil, Right: nil},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, lowestCommonAncestor(item.args, item.p, item.q), "输入:%s", item.args, item.p, item.q)

	}
}
