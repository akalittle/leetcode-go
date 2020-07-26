package LinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//head = [3,2,0,-4], pos = 1


// todo

func Test_hasCircle(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args *ListNode
		want bool
	}{
		{
			args: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  -4,
							Next: nil,
						},
					},
				},
			},
			want: true,
		},
		//{
		//	args: &ListNode{
		//		Val: 1,
		//		Next: &ListNode{
		//			Val:  2,
		//			Next: nil,
		//		},
		//	},
		//	want: true,
		//},
		//{
		//	args: &ListNode{
		//		Val:  1,
		//		Next: nil,
		//	},
		//	want: false,
		//},
	}

	for _, item := range tests {
		ast.Equal(item.want, HasCycle(item.args), "输入:%s", item.args)
	}
}
