package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchRange(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args   []int
		target int
		want   []int
	}{
		{
			args: []int{
				5, 7, 7, 8, 8, 10,
			},
			target: 8,
			want:   []int{3, 4},
		},
		{
			args: []int{
				5, 7, 7, 8, 8, 10,
			},
			target: 6,
			want:   []int{-1, -1},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, searchRange(item.args, item.target), "输入:%s", item.args)
	}
}
