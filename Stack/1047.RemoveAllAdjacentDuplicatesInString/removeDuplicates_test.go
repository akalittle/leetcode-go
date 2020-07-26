package Stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RemoveDuplicates(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args string

		want string
	}{
		{
			args: "abbaca",
			want: "ca",
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, RemoveDuplicates(item.args), "输入:%s", item.args)
	}
}
