package DP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LongestValidParentheses(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args [][]int
		want int
	}{
		{
			args: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, minPathSum(item.args), "输入:%s", item.args)
	}
}
