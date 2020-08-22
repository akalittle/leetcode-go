package DP

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want int
	}{
		{
			args: 1,
			want: 1,
		},
		{
			args: 2,
			want: 2,
		},
		{
			args: 3,
			want: 3,
		},
		{
			args: 4,
			want: 5,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, climbStairs(item.args), "输入:%s", item.args)
	}
}
