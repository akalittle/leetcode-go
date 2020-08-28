package Math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		args int
		want int
	}{
		{
			args: 121,
			want: 121,
		},
		{
			args: -102,
			want: -201,
		},
		{
			args: 0,
			want: 0,
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, reverse(item.args), "输入:%s", item.args)
	}
}
