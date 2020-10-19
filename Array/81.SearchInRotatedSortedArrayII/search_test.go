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
		want   bool
	}{
		{
			args: []int{
				2, 5, 6, 0, 0, 1, 2,
			},
			target: 0,
			want:   true,
		},
		{
			args: []int{
				2, 5, 6, 0, 0, 1, 2,
			},
			target: 3,
			want:   false,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, search(item.args, item.target), "输入:%s", item.args, item.target)
	}
}
