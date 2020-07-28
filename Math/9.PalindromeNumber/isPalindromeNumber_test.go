package Math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindromeNumber(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want bool
	}{
		{
			args: 121,
			want: true,
		},
		{
			args: -102,
			want: false,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, isPalindrome(item.args), "输入:%s", item.args)
	}
}
