package Tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchBST(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		want *TreeNode
	}{
		{
			args: &ListNode{
				//1, 2, 3,
				Val: -10,
				Next: &ListNode{
					Val: -3,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},

			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val:   -10,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, sortedListToBST(item.args), "输入:%s", item.args)

	}
}
