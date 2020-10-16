package Math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	ast := assert.New(t)

	tests := []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}

	for _, item := range tests {
		ast.Equal(item.want, addBinary(item.a, item.b), "输入:%s", item.a, item.b)
	}
}
