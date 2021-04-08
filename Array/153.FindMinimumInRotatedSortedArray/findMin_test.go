package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{
				3, 4, 5, 1, 2,
			},
			want: 1,
		},
		{
			args: []int{
				4, 5, 6, 7, 0, 1, 2,
			},
			want: 0,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, findMin(item.args), "输入:%s", item.args)
	}
}
