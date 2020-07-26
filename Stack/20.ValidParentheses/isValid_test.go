package Stack

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
			args: "()[]{}",
			want: true,
		},
		{
			args: "([)]",
			want: false,
		},
		{
			args: "{[]}",
			want: true,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, isValid(item.args), "输入:%s", item.args)
	}
}
