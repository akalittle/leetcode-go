package LCOF

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want int
	}{
		{
			args: 5,
			want: 5,
		},
		{
			args: 2,
			want: 1,
		},
		{
			args: 6,
			want: 8,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, fib(item.args), "输入:%s", item.args)
		ast.Equal(item.want, fib2Recursive(item.args), "输入:%s", item.args)
	}
}
