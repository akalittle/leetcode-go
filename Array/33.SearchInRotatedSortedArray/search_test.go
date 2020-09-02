package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args   []int
		target int
		want   int
	}{
		{
			args: []int{
				4, 5, 6, 7, 0, 1, 2,
			},
			target: 0,
			want:   4,
		},
		{
			args: []int{
				4, 5, 6, 7, 0, 1, 2,
			},
			target: 3,
			want:   -1,
		},

		{
			args: []int{
				1,
			},
			target: 0,
			want:   -1,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, search(item.args, item.target), "输入:%s", item.args)
	}
}
