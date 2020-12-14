package LCOF

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numWays(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want int
	}{
		{
			args: 7,
			want: 21,
		},
		{
			args: 0,
			want: 1,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, numWays(item.args), "输入:%s", item.args)
	}
}
