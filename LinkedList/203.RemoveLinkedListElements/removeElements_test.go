package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Partition(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		val  int
		want *ListNode
	}{
		{
			args: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  2,
									Next: nil,
								},
							},
						},
					},
				},
			},
			val: 4,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  2,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, RemoveElements(item.args, item.val), "输入:%s", item.args)
		ast.Equal(item.want, RemoveElementsWithRecursive(item.args, item.val), "输入:%s", item.args)
		ast.Equal(item.want, RemoveElementsWithDummyHead(item.args, item.val), "输入:%s", item.args)
	}
}
