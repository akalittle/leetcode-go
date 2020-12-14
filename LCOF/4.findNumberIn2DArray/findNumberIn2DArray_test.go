package LCOF

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findNumberIn2DArray(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args [][]int
		k    int
		want bool
	}{
		{
			args: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			k:    5,
			want: true,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, findNumberIn2DArray(item.args, item.k), "输入:%s", item.args)
	}
}
