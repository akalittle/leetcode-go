package LCOF

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findRepeatNumber(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		want int
	}{
		{
			args: []int{2, 3, 1, 0, 2, 5, 3},
			want: 2,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, findRepeatNumber(item.args), "输入:%s", item.args)
	}
}
