package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_plusOne(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want []int
	}{
		{
			args: []int{
				1, 1, 1, 2, 2, 3,
			},
			want: []int{
				1, 1, 1, 2, 2, 4,
			},
		},
		{
			args: []int{
				9, 9,
			},
			want: []int{
				1, 0, 0,
			},
		},
		{
			args: []int{
				1, 9, 9,
			},
			want: []int{
				2, 0, 0,
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, plusOne(item.args), "输入:%s", item.args)
	}
}
