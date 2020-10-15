package Array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_palindromePartitioning(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		want [][]string
	}{
		{
			args: "aab",
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, partition(item.args), "输入:%s", item.args)
	}
}
