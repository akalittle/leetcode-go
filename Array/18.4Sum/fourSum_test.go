package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fourSum(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args   []int
		target int
		want   [][]int
	}{
		{
			args: []int{
				1, 0, -1, 0, -2, 2,
			},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, fourSum(item.args, item.target), "输入:%s", item.args)
	}
}
