package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_transpose(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args [][]int

		want [][]int
	}{
		{
			args: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, transpose(item.args), "输入:%s", item.args)
	}
}
