package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NextLargerNodes(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		want []int
	}{
		{
			args: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: []int{7, 0, 5, 5, 0},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, NextLargerNodes(item.args), "输入:%s", item.args)
	}
}
