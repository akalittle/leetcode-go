package Math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindromeNumber(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		want int
	}{
		{
			args: "LVIII",
			want: 58,
		},
		{
			args: "III",
			want: 3,
		},
		{
			args: "IV",
			want: 4,
		},
		{
			args: "MCMXCIV",
			want: 1994,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, romanToInt(item.args), "输入:%s", item.args)
	}
}
