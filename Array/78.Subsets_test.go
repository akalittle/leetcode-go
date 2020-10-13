package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subsets(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want [][]int
	}{
		{
			args: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, subsets(item.args), "输入:%s", item.args)
	}
}
