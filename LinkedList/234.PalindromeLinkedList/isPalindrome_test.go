package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		want bool
	}{
		{
			args: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  1,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: true,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, IsPalindrome(item.args), "输入:%s", item.args)
	}
}
