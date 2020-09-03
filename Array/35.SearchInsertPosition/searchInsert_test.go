package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args   []int
		target int
		want   int
	}{
		{
			args: []int{
				1, 3, 5, 6,
			},
			target: 5,
			want:   2,
		},
		{
			args: []int{
				1, 3, 5, 6,
			},
			target: 2,
			want:   1,
		},

		{
			args: []int{
				1, 3, 5, 6,
			},
			target: 7,
			want:   4,
		},
		{
			args: []int{
				1, 3, 5, 6,
			},
			target: 0,
			want:   0,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, searchInsert(item.args, item.target), "输入:%s", item.args)
	}
}
