package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args [][]int

		want [][]int
	}{
		{
			args: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, flipAndInvertImage(item.args), "输入:%s", item.args)
	}
}
