package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		k    int
		want *ListNode
	}{
		{
			args: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			k: 2,
			want: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
		},

		{
			args: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			k: 4,
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, rotateRight(item.args, item.k), "输入:%s", item.args)
	}
}
