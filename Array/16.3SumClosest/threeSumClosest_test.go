package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args   []int
		target int
		want   int
	}{
		{
			args: []int{
				-1, 2, 1, -4,
			},
			target: 1,
			want:   2,
		},
		{
			args: []int{
				0, 0, 0,
			},
			target: 1,
			want:   0,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, threeSumClosest(item.args, item.target), "输入:%s", item.args)
	}
}
