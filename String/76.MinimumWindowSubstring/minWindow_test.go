package String

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minWindow(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string
		s    string
		want string
	}{
		{
			args: "ADOBECODEBANC",
			s:    "ABC",
			want: "BANC",
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, minWindow(item.args, item.s), "输入:%s", item.args)
	}
}
