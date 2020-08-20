package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSum(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want [][]int
	}{
		{
			args: []int{
				-1, 0, 1, 2, -1, -4,
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, threeSum(item.args), "输入:%s", item.args)
	}
}
