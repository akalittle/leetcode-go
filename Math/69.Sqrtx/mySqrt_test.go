package Math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindromeNumber(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want int
	}{
		{
			args: 4,
			want: 2,
		},
		{
			args: 8,
			want: 2,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, mySqrt(item.args), "输入:%s", item.args)
	}
}
