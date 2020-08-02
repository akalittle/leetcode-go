package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumRange(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args []int
		l    int
		r    int
		want int
	}{
		{
			args: []int{
				1, 1, 1, 2, 2, 3,
			},
			l:    0,
			r:    2,
			want: 3,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, Constructor(item.args).SumRange(item.l, item.r), "输入:%s", item.args)
	}
}
