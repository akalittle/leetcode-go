package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string

		want bool
	}{
		{
			args: "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			args: "race a car",
			want: false,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, isPalindrome(item.args), "输入:%s", item.args)
	}
}
