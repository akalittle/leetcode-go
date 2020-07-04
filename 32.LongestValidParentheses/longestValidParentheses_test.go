package _32_LongestValidParentheses_

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LongestValidParentheses(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		want int
	}{
		{
			args: "(())",
			want: 4,
		},
		{
			args: ")()())",
			want: 4,
		},
		{
			args: "(()",
			want: 2,
		},
		{
			args: "(",
			want: 0,
		},
		{
			args: "()(())",
			want: 6,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, LongestValidParentheses(item.args), "输入:%s", item.args)
	}
}
